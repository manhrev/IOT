package data

import (
	"context"
	"time"

	"github.com/manhrev/IOT/server/pkg/ent"
	"github.com/manhrev/IOT/server/pkg/ent/data"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Data interface {
	Create(ctx context.Context, feedName string, value int64, recordedAt *timestamppb.Timestamp) error
	List(ctx context.Context, feedName string) ([]*ent.Data, error)
}

type dataImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Data {
	return &dataImpl{
		entClient: entClient,
	}
}

func (d *dataImpl) Create(ctx context.Context, feedName string, value int64, recordedAt *timestamppb.Timestamp) error {
	recorded := time.Now()
	if recordedAt != nil {
		recorded = recordedAt.AsTime()
	}
	feed, err := d.entClient.Feed.Query().Where(feed.FeedNameEQ(feedName)).Only(ctx)
	if err != nil {
		return err
	}
	_, err = d.entClient.Data.Create().SetFeed(feed).SetData(value).SetRecordedAt(recorded).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *dataImpl) List(ctx context.Context, feedName string) ([]*ent.Data, error) {
	_, err := d.entClient.Feed.Query().Where(feed.FeedNameEQ(feedName)).Only(ctx)
	if err != nil {
		return nil, err
	}
	dat, err := d.entClient.Data.Query().Where(data.HasFeedWith(feed.FeedNameEQ(feedName))).Order(ent.Desc(data.FieldRecordedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
