package service

import (
	"context"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	"github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage/postgres"
	"go.uber.org/zap"
)

type ClientService struct {
	storage *postgres.Storage
	pb.UnimplementedClientServiceServer
}

func NewClientService(storage *postgres.Storage) *ClientService {
	return &ClientService{
		storage: storage,
	}
}

func (s *ClientService) CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.ResponseMessage, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().CreateTender(ctx, req)
	if err != nil {
		log.Error("Error creating tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ClientService) GetMyTenders(ctx context.Context, req *pb.GetMyTendersReq) (*pb.TendersList, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().GetMyTenders(ctx, req)
	if err != nil {
		log.Error("Error fetching my tenders: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ClientService) GetAllTenders(ctx context.Context, req *pb.GetAllTendersReq) (*pb.TendersList, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().GetAllTenders(ctx, req)
	if err != nil {
		log.Error("Error fetching all tenders: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ClientService) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderReq) (*pb.ResponseMessage, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().UpdateTenderStatus(ctx, req)
	if err != nil {
		log.Error("Error updating tender status: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ClientService) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.ResponseMessage, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().DeleteTender(ctx, req)
	if err != nil {
		log.Error("Error deleting tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ClientService) SelectWinner(ctx context.Context, req *pb.SelectWinnerReq) (*pb.ResponseMessage, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Client().SelectWinner(ctx, req)
	if err != nil {
		log.Error("Error deleting tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}
