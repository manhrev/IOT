package feed

import "github.com/manhrev/IOT/server/pkg/ent"

type Feed interface {
}

type feedImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Feed {
	return &feedImpl{
		entClient: entClient,
	}
}
