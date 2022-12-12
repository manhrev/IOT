package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) CreateFeed(ctx context.Context, request *info_pb.CreateFeedRequest) (*info_pb.CreateFeedReply, error) {
	err := s.repository.Feed.Create(ctx, request.GetFeedName())
	if token := s.mqttClient.Subscribe(request.GetFeedName(), 0, nil); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	if err != nil {
		return nil, err
	}
	return &info_pb.CreateFeedReply{}, nil
}
