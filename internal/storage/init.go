package storage

import (
	"context"
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type Storage interface {
	Tender() TenderServiceI
}

type TenderServiceI interface {
	CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.TenderResponse, error)
	GetTenders(ctx context.Context, req *pb.GetTendersReq) (*pb.TendersListResponse, error)
	UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderStatusReq) (*pb.TenderResponse, error)
	DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.TenderResponse, error)
}
