package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *infoServer) GetFeed(ctx context.Context, request *info_pb.GetFeedRequest) (*info_pb.GetFeedReply, error) {
	feedInfo, err := s.repository.Feed.Get(ctx, request.GetFeedName())
	if err != nil {
		return nil, err
	}
	return &info_pb.GetFeedReply{
		Info: &info_pb.FeedInfo{
			FeedId:      int64(feedInfo.ID),
			FeedName:    feedInfo.FeedName,
			CreatedAt:   timestamppb.New(feedInfo.CreatedAt),
			DisplayType: info_pb.FeedDisplayType(feedInfo.DisplayType),
			DataType:    info_pb.FeedDataType(feedInfo.DataType),
		},
	}, nil
}
