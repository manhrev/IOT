package group

import (
	"context"
	"errors"
	"time"

	"github.com/manhrev/IOT/server/pkg/ent"
	"github.com/manhrev/IOT/server/pkg/ent/feed"
	"github.com/manhrev/IOT/server/pkg/ent/group"
)

type Group interface {
	Create(ctx context.Context, groupName string) error
	List(ctx context.Context) ([]*ent.Group, error)
	Delete(ctx context.Context, groupName string) error
	AddFeeds(ctx context.Context, groupName string, feedList []string) error
}

type groupImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Group {
	return &groupImpl{
		entClient: entClient,
	}
}

func (g *groupImpl) Create(ctx context.Context, groupName string) error {
	_, err := g.entClient.Group.Create().SetGroupName(groupName).SetCreatedAt(time.Now()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *groupImpl) List(ctx context.Context) ([]*ent.Group, error) {
	groupList, err := g.entClient.Group.Query().WithFeeds().Order(ent.Desc(group.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return groupList, nil
}

func (g *groupImpl) Delete(ctx context.Context, groupName string) error {

	num_deleted, err := g.entClient.Group.Delete().Where(group.GroupNameEQ(groupName)).Exec(ctx)
	if err != nil {
		return err
	}

	if num_deleted == 0 {
		return errors.New("ent: no records were deleted")
	}

	return nil
}

func (g *groupImpl) AddFeeds(ctx context.Context, groupName string, feedList []string) error {
	group, err := g.entClient.Group.Query().Where(group.GroupNameEQ(groupName)).Only(ctx)
	if err != nil {
		return err
	}

	feeds, err := g.entClient.Feed.Query().Where(feed.FeedNameIn(feedList...)).All(ctx)
	if err != nil {
		return err
	}

	groupUpdate := group.Update()
	for _, feed := range feeds {
		groupUpdate.AddFeeds(feed)
	}

	err = groupUpdate.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
