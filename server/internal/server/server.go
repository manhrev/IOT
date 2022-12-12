package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"

	"github.com/manhrev/IOT/server/internal/server/info"
	m "github.com/manhrev/IOT/server/internal/server/mqtt"
	pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/grpc"
)

const (
	db_user_name  string = "root"
	db_password   string = "password@1"
	db_domain     string = "info_db"
	db_port       string = "3306"
	db_name       string = "info"
	mqtt_host     string = "broker"
	mqtt_port     string = "1883"
	mqtt_clientID string = "server_IOT"
)

func Run() {
	server := newServer()
	Serve(server)
}

func newServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Serve(server *grpc.Server) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

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

	mqttServer := m.New(mqtt_host, mqtt_port, mqtt_clientID, entClient)
	err = mqttServer.Start()

	if err != nil {
		log.Fatalf("Failed when start mqtt Client: %v", err)
	}

	pb.RegisterInfoServer(server, info.NewServer(entClient, mqttServer))

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("error while create listen: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
