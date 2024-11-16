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
	// Build the base query
	query := `SELECT id, client_id, title, description, deadline, budget, status, attachment 
			  FROM tenders 
			  WHERE client_id = $1`

	// Append conditions based on request parameters
	var args []interface{}
	args = append(args, req.ClientId)

	// Filter by budget if provided
	if req.Budget > 0 {
		query += ` AND budget <= $` + fmt.Sprint(len(args)+1)
		args = append(args, req.Budget)
	}

	// Filter by deadline if provided
	if req.Deadline != "" {
		query += ` AND deadline = $` + fmt.Sprint(len(args)+1)
		args = append(args, req.Deadline)
	}

	// Add sorting based on the order parameter
	if req.Order {
		query += ` ORDER BY budget ASC`
	} else {
		query += ` ORDER BY budget DESC`
	}

	// Filter to show only the best tenders if `best` is true
	if req.Best {
		query += ` LIMIT 5` // Show top 5 best tenders, you can adjust as needed
	}

	// Execute the query
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to fetch tenders: %v", err)
		return nil, fmt.Errorf("failed to fetch tenders: %v", err)
	}
	defer rows.Close()

	// Prepare the tenders list
	var tenders []*pb.Tender

	for rows.Next() {
		var tender pb.Tender
		if err := rows.Scan(&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.Attachment); err != nil {
			log.Printf("Failed to scan tender row: %v", err)
			return nil, fmt.Errorf("failed to scan tender row: %v", err)
		}
		tenders = append(tenders, &tender)
	}

	// Return the list of tenders
	return &pb.TendersList{Tenders: tenders}, nil
}

func (c *ClientStorage) GetAllTenders(ctx context.Context, req *pb.GetAllTendersReq) (*pb.TendersList, error) {
	// Build the base query
	query := `SELECT id, client_id, title, description, deadline, budget, status, attachment 
			  FROM tenders WHERE 1=1`

	// Create a slice to hold query arguments
	var args []interface{}

	// Filter by budget if provided
	if req.Budget > 0 {
		query += ` AND budget <= $` + fmt.Sprint(len(args)+1)
		args = append(args, req.Budget)
	}

	// Filter by deadline if provided
	if req.Deadline != "" {
		query += ` AND deadline = $` + fmt.Sprint(len(args)+1)
		args = append(args, req.Deadline)
	}

	// Add sorting based on the order parameter (ascending or descending)
	if req.Order {
		query += ` ORDER BY budget ASC`
	} else {
		query += ` ORDER BY budget DESC`
	}

	// Filter to show only the best tenders if `best` is true
	if req.Best {
		query += ` LIMIT 5` // Show top 5 best tenders (you can adjust this number as needed)
	}

	// Execute the query
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to fetch tenders: %v", err)
		return nil, fmt.Errorf("failed to fetch tenders: %v", err)
	}
	defer rows.Close()

	// Prepare the tenders list
	var tenders []*pb.Tender

	for rows.Next() {
		var tender pb.Tender
		if err := rows.Scan(&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.Attachment); err != nil {
			log.Printf("Failed to scan tender row: %v", err)
			return nil, fmt.Errorf("failed to scan tender row: %v", err)
		}
		tenders = append(tenders, &tender)
	}

	// Return the list of tenders
	return &pb.TendersList{Tenders: tenders}, nil
}
func (c *ClientStorage) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.ResponseMessage, error) {
	// Validate that the tender_id and status are provided
	if req.TenderId == "" || req.Status == "" {
		return nil, fmt.Errorf("tender_id and status are required")
	}

	// Build the SQL query to update the tender status
	query := `UPDATE tenders SET status = $1, description = $2, budget = $3 WHERE id = $4`

	// Execute the query with the provided parameters
	_, err := c.db.ExecContext(ctx, query, req.Status, req.Description, req.Budget, req.TenderId)
	if err != nil {
		log.Printf("Failed to update tender status: %v", err)
		return nil, fmt.Errorf("failed to update tender status: %v", err)
	}

	// Return success response
	return &pb.ResponseMessage{
		Message: "Tender status updated successfully",
	}, nil
}
func (c *ClientStorage) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.ResponseMessage, error) {
	// Validate that the tender_id is provided
	if req.TenderId == "" {
		return nil, fmt.Errorf("tender_id is required")
	}

	// Build the SQL query to delete the tender by tender_id
	query := `DELETE FROM tenders WHERE id = $1`

	// Execute the query with the provided tender_id
	result, err := c.db.ExecContext(ctx, query, req.TenderId)
	if err != nil {
		log.Printf("Failed to delete tender: %v", err)
		return nil, fmt.Errorf("failed to delete tender: %v", err)
	}

	// Check if any rows were affected (i.e., a tender was deleted)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		// No rows were affected, meaning the tender was not found
		return &pb.ResponseMessage{
			Message: "Tender not found or already deleted",
		}, nil
	}

	// Return success response
	return &pb.ResponseMessage{
		Message: "Tender deleted successfully",
	}, nil
}
