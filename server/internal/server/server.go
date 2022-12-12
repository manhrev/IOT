package server

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manhrev/IOT/server/internal/server/info"
	pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/grpc"
)

const (
	db_user_name string = "root"
	db_password  string = "password@1"
	db_domain    string = "info_db"
	db_port      string = "3306"
	db_name      string = "info"
)

func Run() {
	server := newServer()
	Serve(server)
}

func newServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}

func Serve(server *grpc.Server) {
	entClient, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", db_user_name, db_password, db_domain, db_port, db_name))
	if err != nil {
		log.Fatalf("cannot create database connection: %v", err)
	}
	defer func() {
		if err := entClient.Close(); err != nil {
			panic(err)
		}
	}()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	pb.RegisterInfoServer(server, info.NewServer(entClient))

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("error while create listen: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
