package postgres

import (
	"context"
	"database/sql"

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
	return nil, nil
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
