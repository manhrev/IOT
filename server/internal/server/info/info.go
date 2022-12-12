package info

import (
	"github.com/manhrev/IOT/server/internal/repository"
	m "github.com/manhrev/IOT/server/internal/server/mqtt"
	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
)

func NewServer(entClient *ent.Client, mqttClient *m.MqttServer) info_pb.InfoServer {
	return &infoServer{
		repository: repository.New(entClient),
		mqttClient: mqttClient,
	}
}

type infoServer struct {
	info_pb.UnimplementedInfoServer
	repository *repository.Repository
	mqttClient *m.MqttServer
}
