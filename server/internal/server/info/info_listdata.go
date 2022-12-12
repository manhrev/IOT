package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *infoServer) ListData(ctx context.Context, request *info_pb.ListDataRequest) (*info_pb.ListDataReply, error) {
	data, err := s.repository.Data.List(ctx, request.GetFeedName())
	if err != nil {
		return nil, err
	}
	return &info_pb.ListDataReply{
		ListData: convertEntDataToData(data),
	}, nil
}

func convertEntDataToData(entData []*ent.Data) []*info_pb.Data {
	datapb := []*info_pb.Data{}
	for _, entDat := range entData {
		dat := &info_pb.Data{
			Value:      int64(entDat.Data),
			RecordedAt: timestamppb.New(entDat.RecordedAt),
		}
		datapb = append(datapb, dat)
	}
	return datapb
}
