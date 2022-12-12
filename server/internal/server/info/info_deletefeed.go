package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) DeleteFeed(ctx context.Context, request *info_pb.DeleteFeedRequest) (*info_pb.DeleteFeedReply, error) {
	err := s.repository.Feed.Delete(ctx, request.GetFeedName())
	if err != nil {
		return nil, err
	}
	return &info_pb.DeleteFeedReply{}, nil
}
