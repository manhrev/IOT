package repository

import (
	"github.com/manhrev/IOT/server/internal/repository/data"
	"github.com/manhrev/IOT/server/internal/repository/feed"
	"github.com/manhrev/IOT/server/pkg/ent"
)

type Repository struct {
	entClient *ent.Client

	Feed feed.Feed
	Data data.Data
}

func New(entClient *ent.Client) *Repository {
	return &Repository{
		entClient: entClient,
		Feed:      feed.New(entClient),
		Data:      feed.New(entClient),
	}
}
