package info

import (
	"github.com/manhrev/IOT/server/internal/repository"
	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
)

func NewServer(entClient *ent.Client) info_pb.InfoServer {
	return &infoServer{
		repository: repository.New(entClient),
	}
}

type infoServer struct {
	info_pb.UnimplementedInfoServer
	repository *repository.Repository
}
