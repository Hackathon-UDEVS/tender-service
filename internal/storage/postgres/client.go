package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	// Query the database to retrieve tenders based on client_id and status
	query := "SELECT id, client_id, title, description, deadline, budget, status, attachment FROM tenders WHERE client_id = $1"

	// Apply filter for status if provided
	if req.Status != "" {
		query += " AND status = $2"
	}

	rows, err := c.db.QueryContext(ctx, query, req.ClientId, req.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []*pb.Tender
	for rows.Next() {
		var tender pb.Tender
		if err := rows.Scan(&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.Attachment); err != nil {
			return nil, err
		}
		tenders = append(tenders, &tender)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.TendersList{Tenders: tenders}, nil
}

func (c *ClientStorage) GetAllTenders(ctx context.Context, req *pb.GetAllTendersReq) (*pb.TendersList, error) {
	// Query the database to retrieve all tenders
	query := "SELECT id, client_id, title, description, deadline, budget, status, attachment FROM tenders"

	// Apply filter for status if provided
	if req.Status != "" {
		query += " WHERE status = $1"
	}

	rows, err := c.db.QueryContext(ctx, query, req.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []*pb.Tender
	for rows.Next() {
		var tender pb.Tender
		if err := rows.Scan(&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.Attachment); err != nil {
			return nil, err
		}
		tenders = append(tenders, &tender)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.TendersList{Tenders: tenders}, nil
}

func (c *ClientStorage) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderReq) (*pb.ResponseMessage, error) {
	// Prepare the update query with optional fields
	query := "UPDATE tenders SET "
	args := []interface{}{}
	setClauses := []string{}

	if req.Title != "" {
		setClauses = append(setClauses, "title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Title)
	}
	if req.Description != "" {
		setClauses = append(setClauses, "description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Description)
	}
	if req.Deadline != "" {
		setClauses = append(setClauses, "deadline = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Deadline)
	}
	if req.Status != "" {
		setClauses = append(setClauses, "status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Status)
	}
	if req.Budget != 0 {
		setClauses = append(setClauses, "budget = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Budget)
	}
	if req.Attachment != "" {
		setClauses = append(setClauses, "attachment = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Attachment)
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(setClauses, ", ") + " WHERE id = $" + strconv.Itoa(len(args)+1)
	args = append(args, req.TenderId)

	// Execute the query
	_, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update tender: %v", err)
	}

	return &pb.ResponseMessage{
		Message: "Tender updated successfully",
	}, nil
}

func (c *ClientStorage) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.ResponseMessage, error) {
	// Check if the tender exists and belongs to the given client
	var clientID string
	query := "SELECT client_id FROM tenders WHERE id = $1"
	err := c.db.QueryRowContext(ctx, query, req.TenderId).Scan(&clientID)

	// Handle case where tender is not found
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("tender not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query tender: %v", err)
	}

	// Check if the provided client_id matches the tender owner
	if clientID != req.ClientId {
		return nil, fmt.Errorf("you are not authorized to delete this tender")
	}

	// Proceed to delete the tender
	deleteQuery := "DELETE FROM tenders WHERE id = $1"
	_, err = c.db.ExecContext(ctx, deleteQuery, req.TenderId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete tender: %v", err)
	}

	return &pb.ResponseMessage{
		Message: "Tender deleted successfully",
	}, nil
}

func (c *ClientStorage) SelectWinner(ctx context.Context, req *pb.SelectWinnerReq) (*pb.ResponseMessage, error) {
	// Convert client_id to int (since req.ClientId is a string)
	clientID, err := strconv.Atoi(req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %v", err)
	}

	// Step 1: Verify if the tender belongs to the provided client
	var tenderClientID int
	var tenderStatus string
	query := `SELECT client_id, status FROM tenders WHERE id = $1`
	err = c.db.QueryRowContext(ctx, query, req.TenderId).Scan(&tenderClientID, &tenderStatus)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("tender not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query tender: %v", err)
	}

	// Ensure the client is the owner of the tender
	if tenderClientID != clientID {
		return nil, fmt.Errorf("you are not authorized to select the winner for this tender")
	}

	// Step 2: Ensure the tender status is "open" before selecting a winner
	if tenderStatus != "open" {
		return nil, fmt.Errorf("the tender is not open for selection")
	}

	// Step 3: Update all bids to "rejected" (except for the selected one)
	_, err = c.db.ExecContext(ctx, `
        UPDATE bid
        SET status = 'rejected'
        WHERE tender_id = $1 AND status != 'accepted'`, req.TenderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update bid statuses: %v", err)
	}

	// Step 4: Set the selected contractor's status to "accepted"
	_, err = c.db.ExecContext(ctx, `
        UPDATE bid
        SET status = 'accepted'
        WHERE tender_id = $1 AND contractor_id = $2`, req.TenderId, req.ContractorId)
	if err != nil {
		return nil, fmt.Errorf("failed to update selected contractor's status: %v", err)
	}

	// Step 5: Update the tender status to "awarded"
	_, err = c.db.ExecContext(ctx, `
        UPDATE tenders
        SET status = 'awarded'
        WHERE id = $1`, req.TenderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update tender status: %v", err)
	}

	// Step 6: Optionally, you can send notifications or handle other business logic here

	// Return success message
	return &pb.ResponseMessage{
		Message: "Winner selected and tender awarded successfully",
	}, nil
}
