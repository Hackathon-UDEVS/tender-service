package service

import (
	"context"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	"github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage/postgres"
	"go.uber.org/zap"
)

type ContractorService struct {
	storage *postgres.Storage
	pb.UnimplementedContractorServiceServer
}

func NewContractorService(storage *postgres.Storage) *ContractorService {
	return &ContractorService{
		storage: storage,
	}
}

func (s *ContractorService) SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Contractor().SubmitBid(ctx, req)
	if err != nil {
		log.Error("Error submitting bid: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *ContractorService) GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	log, _ := logger.NewLogger()
	res, err := s.storage.Contractor().GetBidsForTender(ctx, req)
	if err != nil {
		log.Error("Error fetching bids for tender: ", zap.Error(err))
		return nil, err
	}
	return res, nil
}
