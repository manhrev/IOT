package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) CreateGroup(ctx context.Context, request *info_pb.CreateGroupRequest) (*info_pb.CreateGroupReply, error) {
	err := s.repository.Group.Create(ctx, request.GetGroupName())

	if err != nil {
		return nil, err
	}
	return &info_pb.CreateGroupReply{}, nil
}
