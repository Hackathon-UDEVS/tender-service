package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
	// Validate input
	if req.TenderId == "" || req.ContractorId == "" || req.Price <= 0 || req.DeliveryTime == "" {
		return nil, fmt.Errorf("invalid input, all fields are required")
	}

	// Prepare the SQL query to insert the bid
	query := `INSERT INTO bids (tender_id, contractor_id, price, delivery_time, comments, status)
			  VALUES ($1, $2, $3, $4, $5, 'submitted')`

	// Execute the query
	_, err := c.db.ExecContext(ctx, query, req.TenderId, req.ContractorId, req.Price, req.DeliveryTime, req.Comments)
	if err != nil {
		log.Printf("Failed to submit bid: %v", err)
		return nil, fmt.Errorf("failed to submit bid: %v", err)
	}

	// Return a success response
	return &pb.BidResponse{
		Message: "Bid submitted successfully",
	}, nil
}

// GetBidsForTender retrieves all bids for a specific tender, with optional filters
func (c *ContractorStorage) GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	// Validate input
	if req.TenderId == "" {
		return nil, fmt.Errorf("tender_id is required")
	}

	// Build the SQL query with optional filters
	query := `SELECT id, tender_id, contractor_id, price, delivery_time, comments, status 
			  FROM bids WHERE tender_id = $1`

	var args []interface{}
	args = append(args, req.TenderId)

	if req.Budget > 0 {
		query += ` AND price <= $2`
		args = append(args, req.Budget)
	}

	if req.Deadline != "" {
		// You can implement filtering by deadline here if needed
		// query += ` AND delivery_time <= $3`
		// args = append(args, req.Deadline)
	}

	if req.Order {
		query += ` ORDER BY price ASC` // Sort by price (ascending)
	} else {
		query += ` ORDER BY price DESC` // Sort by price (descending)
	}

	if req.Best {
		query += ` LIMIT 1` // Only return the best bid
	}

	// Execute the query
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to get bids: %v", err)
		return nil, fmt.Errorf("failed to get bids: %v", err)
	}
	defer rows.Close()

	// Collect the results
	var bids []*pb.Bid
	for rows.Next() {
		var bid pb.Bid
		if err := rows.Scan(&bid.Id, &bid.TenderId, &bid.ContractorId, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status); err != nil {
			log.Printf("Failed to scan bid: %v", err)
			return nil, fmt.Errorf("failed to scan bid: %v", err)
		}
		bids = append(bids, &bid)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		log.Printf("Failed to iterate rows: %v", err)
		return nil, fmt.Errorf("failed to iterate rows: %v", err)
	}

	// Return the list of bids
	return &pb.BidsListResponse{
		Bids: bids,
	}, nil
}
