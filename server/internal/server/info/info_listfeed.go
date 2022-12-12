package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) ListFeed(ctx context.Context, request *info_pb.ListFeedRequest) (*info_pb.ListFeedReply, error) {
	return &info_pb.ListFeedReply{}, nil
}
