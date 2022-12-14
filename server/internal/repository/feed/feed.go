package feed

import (
	"context"
	"errors"
	"time"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"github.com/manhrev/IOT/server/pkg/ent/data"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
)

type Feed interface {
	Create(ctx context.Context, feedName string, dataType info_pb.FeedDataType, displayType info_pb.FeedDisplayType) error
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

func (f *feedImpl) Create(ctx context.Context, feedName string, dataType info_pb.FeedDataType, displayType info_pb.FeedDisplayType) error {
	_, err := f.entClient.Feed.Create().
		SetFeedName(feedName).
		SetDataType(uint16(dataType)).
		SetDisplayType(uint16(displayType)).
		SetCreatedAt(time.Now()).Save(ctx)
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
	feedList, err := f.entClient.Feed.Query().Order(ent.Desc(feed.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return feedList, nil
}

func (f *feedImpl) Delete(ctx context.Context, feedName string) error {
	_, err := f.entClient.Data.Delete().Where(data.HasFeedWith(feed.FeedNameEQ(feedName))).Exec(ctx)
	if err != nil {
		return err
	}

	num_deleted, err := f.entClient.Feed.Delete().Where(feed.FeedNameEQ(feedName)).Exec(ctx)
	if err != nil {
		return err
	}

	if num_deleted == 0 {
		return errors.New("ent: no records were deleted")
	}

	return nil
}
