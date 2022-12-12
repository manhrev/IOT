package info

import (
	"context"
	"encoding/json"
	"errors"

	m "github.com/manhrev/IOT/server/internal/server/mqtt"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
)

func (s *infoServer) CreateData(ctx context.Context, request *info_pb.CreateDataRequest) (*info_pb.CreateDataReply, error) {
	payload := m.Payload{
		Value: int(request.GetValue()),
	}
	message, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("Payload is not match with the right type")
	}

	token := s.mqttClient.Publish(request.GetFeedName(), 0, false, message)
	token.Wait()

	if err != nil {
		return nil, err
	}
	return &info_pb.CreateDataReply{}, nil
}
