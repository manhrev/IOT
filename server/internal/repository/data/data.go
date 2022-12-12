package data

import "github.com/manhrev/IOT/server/pkg/ent"

type Data interface {
}

type dataImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Data {
	return &dataImpl{
		entClient: entClient,
	}
}
