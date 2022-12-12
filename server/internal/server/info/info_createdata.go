package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) CreateData(ctx context.Context, request *info_pb.CreateDataRequest) (*info_pb.CreateDataReply, error) {
	err := s.repository.Data.Create(ctx, request.GetFeedName(), request.GetValue(), request.GetRecordedAt())
	if err != nil {
		return nil, err
	}
	return &info_pb.CreateDataReply{}, nil
}
