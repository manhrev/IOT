package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) AddFeedsToGroup(ctx context.Context, request *info_pb.AddFeedsToGroupRequest) (*info_pb.AddFeedsToGroupReply, error) {
	err := s.repository.Group.AddFeeds(ctx, request.GetGroupName(), request.GetFeedNames())
	if err != nil {
		return nil, err
	}
	return &info_pb.AddFeedsToGroupReply{}, nil
}
