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

// SubmitBid inserts a new bid into the database
func (c *ContractorStorage) SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error) {
	return nil, nil
}

// GetBidsForTender retrieves all bids for a specific tender, with optional filters
func (c *ContractorStorage) GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	return nil, nil
}

// Gets only bids of mine
func (c *ContractorStorage) GetMyBids(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	return nil, nil
}
