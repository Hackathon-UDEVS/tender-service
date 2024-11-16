package postgres

import (
	"context"
	"database/sql"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type ContractorStorage struct {
	db *sql.DB
}

func NewContractorStorage(db *sql.DB) *ContractorStorage {
	return &ContractorStorage{
		db: db,
	}
}

func (c *ContractorStorage) SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error) {
	return nil, nil
}
func (c *ContractorStorage) GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	return nil, nil
}
