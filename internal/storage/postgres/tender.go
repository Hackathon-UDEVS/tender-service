package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	"github.com/go-redis/redis"
	"github.com/lib/pq"
)

type TenderRepo struct {
	Db *sql.DB
	Rd *redis.Client
}

func NewTenderRepo(db *sql.DB, rd *redis.Client) *TenderRepo {
	return &TenderRepo{
		Db: db,
		Rd: rd,
	}
}

func (d *TenderRepo) CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.TenderResponse, error) {
	query := `
		INSERT INTO tender (client_id, title, description, attachment, deadline, budget, status)
		VALUES ($1, $2, $3, $4, $5, $6, 'open') RETURNING id;
	`
	var tenderID int
	deadline, err := time.Parse("2006-01-02 15:04", req.Deadline)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format: %w", err)
	}

	err = d.Db.QueryRowContext(ctx, query,
		req.ClientId,
		req.Title,
		req.Description,
		req.Attachment,
		deadline,
		req.Budget,
	).Scan(&tenderID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf("database error: %s", pqErr.Message)
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	cacheKey := fmt.Sprintf("tender:%d", tenderID)
	err = d.Rd.Set(cacheKey, tenderID, 24*time.Hour).Err()
	if err != nil {
		fmt.Printf("warning: failed to cache tender ID: %v\n", err)
	}

	return &pb.TenderResponse{
		Message: fmt.Sprintf("Tender created successfully with ID %d", tenderID),
	}, nil
}

func (d *TenderRepo) GetTenders(ctx context.Context, req *pb.GetTendersReq) (*pb.TendersListResponse, error) {
	query := `
		SELECT id, client_id, title, description, attachment, deadline, budget, status
		FROM tender WHERE client_id = $1;
	`

	rows, err := d.Db.QueryContext(ctx, query, req.ClientId)
	if err != nil {
		return nil, fmt.Errorf("failed to query tenders: %w", err)
	}
	defer rows.Close()

	var tenders []*pb.Tender
	for rows.Next() {
		var tender pb.Tender
		var deadline time.Time

		err = rows.Scan(
			&tender.Id,
			&tender.ClientId,
			&tender.Title,
			&tender.Description,
			&tender.Attachment,
			&deadline,
			&tender.Budget,
			&tender.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tender: %w", err)
		}

		tender.Deadline = deadline.Format("2006-01-02 15:04")
		tenders = append(tenders, &tender)
	}

	return &pb.TendersListResponse{Tenders: tenders}, nil
}

func (d *TenderRepo) UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.TenderResponse, error) {
	query := `
		UPDATE tender
		SET status = $1, description = $2, budget = $3
		WHERE id = $4 RETURNING id;
	`

	_, err := d.Db.ExecContext(ctx, query, req.Status, req.Description, req.Budget, req.TenderId)
	if err != nil {
		return nil, fmt.Errorf("failed to update tender: %w", err)
	}

	return &pb.TenderResponse{
		Message: fmt.Sprintf("Tender status updated successfully with ID %s", req.TenderId),
	}, nil
}

func (d *TenderRepo) DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.TenderResponse, error) {
	query := `
		DELETE FROM tender WHERE id = $1 RETURNING id;
	`

	var deletedID int
	err := d.Db.QueryRowContext(ctx, query, req.TenderId).Scan(&deletedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tender not found")
		}
		return nil, fmt.Errorf("failed to delete tender: %w", err)
	}

	cacheKey := fmt.Sprintf("tender:%d", deletedID)
	err = d.Rd.Del(cacheKey).Err()
	if err != nil {
		fmt.Printf("warning: failed to remove tender from cache: %v\n", err)
	}

	return &pb.TenderResponse{
		Message: fmt.Sprintf("Tender deleted successfully with ID %d", deletedID),
	}, nil
}
