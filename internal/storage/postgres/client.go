package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type ClientStorage struct {
	db *sql.DB
}

func NewClientStorage(db *sql.DB) *ClientStorage {
	return &ClientStorage{
		db: db,
	}
}

func (c *ClientStorage) CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.ResponseMessage, error) {
	// Prepare SQL statement to insert the new tender into the database
	query := `INSERT INTO tenders (client_id, title, description, deadline, budget, attachment)
			  VALUES ($1, $2, $3, $4, $5, $6)`

	// Execute the query
	_, err := c.db.ExecContext(ctx, query, req.ClientId, req.Title, req.Description, req.Deadline, req.Budget, req.Attachment)
	if err != nil {
		log.Printf("Failed to create tender: %v", err)
		return nil, fmt.Errorf("failed to create tender: %v", err)
	}

	// Prepare response message
	response := &pb.ResponseMessage{
		Message: "Tender created successfully",
	}

	return response, nil
}

func (c *ClientStorage) GetMyTenders(ctx context.Context, req *pb.GetMyTendersReq) (*pb.TendersList, error) {
	return nil, nil
}

func (c *ClientStorage) GetAllTenders(ctx context.Context, req *pb.GetAllTendersReq) (*pb.TendersList, error) {
	return nil, nil
}

func (c *ClientStorage) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.ResponseMessage, error) {
	return nil, nil
}

func (c *ClientStorage) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.ResponseMessage, error) {
	return nil, nil
}
