package data

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manhrev/IOT/server/pkg/ent"
)

const (
	_driver = "mysql"
	_url    = "root:password@1@tcp(localhost:33306)/info?charset=utf8&parseTime=true"
)

func Test_List(t *testing.T) {
	entClient, err := ent.Open(_driver, _url)
	if err != nil {
		log.Fatalf("error creating connection to database%v", err.Error())
	}
	var (
		ctx    = context.Background()
		client = entClient
	)

	dataRepo := New(client)
	err = dataRepo.Create(
		ctx,
		"led",
		12,
		nil,
	)
	if err != nil {
		log.Fatalf("error while querying:%s", err.Error())
	}

}
