package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sort"

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
	// Validate tender existence
	var tenderID string
	query := `SELECT id FROM tenders WHERE id = $1`
	err := c.db.QueryRowContext(ctx, query, req.TenderId).Scan(&tenderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tender not found")
		}
		return nil, fmt.Errorf("failed to query tender: %v", err)
	}

	// Insert the new bid into the database
	insertQuery := `
		INSERT INTO bid (tender_id, contractor_id, price, delivery_time, comments, status)
		VALUES ($1, $2, $3, $4, $5, 'submitted')`
	_, err = c.db.ExecContext(ctx, insertQuery, req.TenderId, req.ContractorId, req.Price, req.DeliveryTime, req.Comments)
	if err != nil {
		return nil, fmt.Errorf("failed to insert bid: %v", err)
	}

	// Return the response
	return &pb.BidResponse{
		Message: "Bid submitted successfully",
	}, nil
}

// GetBidsForTender retrieves all bids for a specific tender, with optional filters
func (c *ContractorStorage) GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	// Validate tender existence
	var tenderID string
	query := `SELECT id FROM tenders WHERE id = $1`
	err := c.db.QueryRowContext(ctx, query, req.TenderId).Scan(&tenderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tender not found")
		}
		return nil, fmt.Errorf("failed to query tender: %v", err)
	}

	// Query all bids for the tender with optional filters
	var bids []*pb.Bid // Change to slice of pointers
	selectQuery := `SELECT id, tender_id, contractor_id, price, delivery_time, comments, status FROM bid WHERE tender_id = $1 AND status = $2`
	rows, err := c.db.QueryContext(ctx, selectQuery, req.TenderId, req.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to query bids: %v", err)
	}
	defer rows.Close()

	// Read all bids from the database and append them as pointers
	for rows.Next() {
		var bid pb.Bid
		err := rows.Scan(&bid.Id, &bid.TenderId, &bid.ContractorId, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bid: %v", err)
		}
		// Append pointer to the bid
		bids = append(bids, &bid)
	}

	// Function to convert delivery time string (e.g., "2 weeks", "5 days") to an integer (number of days)
	convertToDays := func(deliveryTime string) (int, error) {
		var days int
		var unit string
		// Assuming the format is like "2 weeks", "5 days", etc.
		_, err := fmt.Sscanf(deliveryTime, "%d %s", &days, &unit)
		if err != nil {
			return 0, fmt.Errorf("invalid delivery time format: %v", err)
		}

		// Convert to days (assuming 1 week = 7 days)
		switch unit {
		case "week", "weeks":
			days *= 7
		case "day", "days":
			// Already in days, no conversion needed
		default:
			return 0, fmt.Errorf("unsupported unit for delivery time: %s", unit)
		}
		return days, nil
	}

	// If 'best' flag is true, calculate the best bids using a weighted formula
	if req.Best {
		// Define weights for price and delivery time
		priceWeight := 0.7 // Price weight (70%)
		timeWeight := 0.3  // Delivery time weight (30%)

		// Sort bids by price and delivery time using a weighted formula
		sort.SliceStable(bids, func(i, j int) bool {
			// Convert delivery time from string to integer (days)
			deliveryTimeI, err := convertToDays(bids[i].DeliveryTime)
			if err != nil {
				log.Printf("Error converting delivery time for bid %s: %v", bids[i].Id, err)
				return false
			}

			deliveryTimeJ, err := convertToDays(bids[j].DeliveryTime)
			if err != nil {
				log.Printf("Error converting delivery time for bid %s: %v", bids[j].Id, err)
				return false
			}

			// Calculate score for both bids
			scoreI := (priceWeight * bids[i].Price) + (timeWeight * float64(deliveryTimeI))
			scoreJ := (priceWeight * bids[j].Price) + (timeWeight * float64(deliveryTimeJ))

			// Lower score is better
			return scoreI < scoreJ
		})

		// Limit to the top 5 bids
		if len(bids) > 5 {
			bids = bids[:5]
		}
	}

	// Return the list of bids as a pointer slice
	return &pb.BidsListResponse{
		Bids: bids, // Now returning a slice of pointers
	}, nil
}

// GetMyBids retrieves all bids submitted by a specific contractor for a specific tender
func (c *ContractorStorage) GetMyBids(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error) {
	// Assuming ContractorId is stored in the context
	contractorID, ok := ctx.Value("contractor_id").(string)
	if !ok {
		return nil, fmt.Errorf("contractor_id not found in context")
	}

	// Validate tender existence
	var tenderID string
	query := `SELECT id FROM tenders WHERE id = $1`
	err := c.db.QueryRowContext(ctx, query, req.TenderId).Scan(&tenderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tender not found")
		}
		return nil, fmt.Errorf("failed to query tender: %v", err)
	}

	// Prepare the query to fetch bids based on the provided parameters
	selectQuery := `
		SELECT id, tender_id, contractor_id, price, delivery_time, comments, status
		FROM bid
		WHERE tender_id = $1 AND contractor_id = $2
	`

	// Add optional filters if they are provided
	var args []interface{}
	args = append(args, req.TenderId, contractorID)

	if req.Budget > 0 {
		selectQuery += " AND price <= $3"
		args = append(args, req.Budget)
	}

	if req.Deadline != "" {
		selectQuery += " AND delivery_time <= $4"
		args = append(args, req.Deadline)
	}

	if req.Status != "" {
		selectQuery += " AND status = $5"
		args = append(args, req.Status)
	}

	// Execute the query
	rows, err := c.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query bids: %v", err)
	}
	defer rows.Close()

	// Prepare the slice of bids
	var bids []*pb.Bid

	// Read all bids from the database
	for rows.Next() {
		var bid pb.Bid
		err := rows.Scan(&bid.Id, &bid.TenderId, &bid.ContractorId, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bid: %v", err)
		}
		bids = append(bids, &bid)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bids: %v", err)
	}

	// Return the response with the list of bids
	return &pb.BidsListResponse{
		Bids: bids,
	}, nil
}
