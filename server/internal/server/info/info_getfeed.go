package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) GetFeed(ctx context.Context, request *info_pb.GetFeedRequest) (*info_pb.GetFeedReply, error) {
	return &info_pb.GetFeedReply{}, nil
}
