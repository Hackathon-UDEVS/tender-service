package service

import (
	pb "github.com/Hackaton-UDEVS/first/internal/genproto/first-service"
	"github.com/Hackaton-UDEVS/first/internal/logger"
	"github.com/Hackaton-UDEVS/first/internal/storage/postgres"
)

type App1Service struct {
	storage *postgres.Storage
	pb.UnimplementedFirstServiceServer
}

func NewApp1Service(storage *postgres.Storage) *App1Service {
	return &App1Service{
		storage: storage,
	}
}

func (s App1Service) TestCreate(req *pb.Test1Req) (*pb.Test1Res, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.App1.TestCreate(req)
	if err != nil {
		log.Error("Error ")
		return nil, err
	}
	return res, nil
}
