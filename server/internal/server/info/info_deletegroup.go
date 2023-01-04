package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) DeleteGroup(ctx context.Context, request *info_pb.DeleteGroupRequest) (*info_pb.DeleteGroupReply, error) {
	err := s.repository.Group.Delete(ctx, request.GetGroupName())
	if err != nil {
		return nil, err
	}
	// Unscribe
	// if token := s.mqttClient.Unsubscribe(request.GetFeedName()); token.Wait() && token.Error() != nil {
	// 	return nil, token.Error()
	// }
	return &info_pb.DeleteGroupReply{}, nil
}
