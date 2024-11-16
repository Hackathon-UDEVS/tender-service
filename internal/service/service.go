package service

import (
	"context"
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	"github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage/postgres"
	"go.uber.org/zap"
)

type TenderServiceImpl struct {
	storage *postgres.Storage
	pb.UnimplementedTenderServiceServer
}

func NewTenderService(storage *postgres.Storage) *TenderServiceImpl {
	return &TenderServiceImpl{
		storage: storage,
	}
}

func (s *TenderServiceImpl) CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Tender.CreateTender(ctx, req)
	if err != nil {
		log.Error("Error creating tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) GetTenders(ctx context.Context, req *pb.GetTendersReq) (*pb.TendersListResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Tender.GetTenders(ctx, req)
	if err != nil {
		log.Error("Error fetching tenders: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Tender.UpdateTenderStatus(ctx, req)
	if err != nil {
		log.Error("Error updating tender status: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Tender.DeleteTender(ctx, req)
	if err != nil {
		log.Error("Error deleting tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}
