package service

import (
	"context"
	"github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage"
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type TenderServiceImpl struct {
	storage storage.Storage
	pb.UnimplementedTenderServiceServer
}

func NewTenderService(storage storage.Storage) *TenderServiceImpl {
	return &TenderServiceImpl{
		storage: storage,
	}
}

func (s *TenderServiceImpl) CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Tender.(ctx, req)
	if err != nil {
		log.Error("Error creating tender: ", err)
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) GetTenders(ctx context.Context, req *pb.GetTendersReq) (*pb.TendersListResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.GetTenders(ctx, req)
	if err != nil {
		log.Error("Error fetching tenders: ", err)
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.UpdateTenderStatus(ctx, req)
	if err != nil {
		log.Error("Error updating tender status: ", err)
		return nil, err
	}
	return res, nil
}

func (s *TenderServiceImpl) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.TenderResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.DeleteTender(ctx, req)
	if err != nil {
		log.Error("Error deleting tender: ", err)
		return nil, err
	}
	return res, nil
}
