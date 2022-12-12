package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) CreateData(ctx context.Context, request *info_pb.CreateDataRequest) (*info_pb.CreateDataReply, error) {
	return &info_pb.CreateDataReply{}, nil
}
