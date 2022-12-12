package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *infoServer) ListFeed(ctx context.Context, request *info_pb.ListFeedRequest) (*info_pb.ListFeedReply, error) {
	feedList, err := s.repository.Feed.List(ctx)
	if err != nil {
		return nil, err
	}

	return &info_pb.ListFeedReply{
		Feeds: convertListEntFeedInfoToFeedInfo(feedList),
	}, nil
}

func convertListEntFeedInfoToFeedInfo(entFeedList []*ent.Feed) []*info_pb.FeedInfo {
	listFeedInfo := []*info_pb.FeedInfo{}
	for _, entFeed := range entFeedList {
		feedInfo := &info_pb.FeedInfo{
			FeedId:    int64(entFeed.ID),
			FeedName:  entFeed.FeedName,
			CreatedAt: timestamppb.New(entFeed.CreatedAt),
		}
		listFeedInfo = append(listFeedInfo, feedInfo)
	}
	return listFeedInfo
}
