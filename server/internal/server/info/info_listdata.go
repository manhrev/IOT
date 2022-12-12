package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) ListData(ctx context.Context, request *info_pb.ListDataRequest) (*info_pb.ListDataReply, error) {
	return &info_pb.ListDataReply{}, nil
}
