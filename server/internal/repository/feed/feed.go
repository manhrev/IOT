package feed

import (
	"context"
	"errors"
	"time"

	"github.com/manhrev/IOT/server/pkg/ent"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
)

type Feed interface {
	Create(ctx context.Context, feedName string) error
	Get(ctx context.Context, feedName string) (*ent.Feed, error)
	List(ctx context.Context) ([]*ent.Feed, error)
	Delete(ctx context.Context, feedName string) error
}

type feedImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Feed {
	return &feedImpl{
		entClient: entClient,
	}
}

func (f *feedImpl) Create(ctx context.Context, feedName string) error {
	_, err := f.entClient.Feed.Create().SetFeedName(feedName).SetCreatedAt(time.Now()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f *feedImpl) Get(ctx context.Context, feedName string) (*ent.Feed, error) {
	feedResult, err := f.entClient.Feed.Query().Where(feed.FeedNameEQ(feedName)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return feedResult, nil
}

func (f *feedImpl) List(ctx context.Context) ([]*ent.Feed, error) {
	feedList, err := f.entClient.Feed.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return feedList, nil
}

func (f *feedImpl) Delete(ctx context.Context, feedName string) error {
	num_deleted, err := f.entClient.Feed.Delete().Where(feed.FeedNameEQ(feedName)).Exec(ctx)
	if err != nil {
		return err
	}

	if num_deleted == 0 {
		return errors.New("ent: no records were deleted")
	}

	return nil
}
