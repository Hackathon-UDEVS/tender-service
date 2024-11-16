package storage

import (
	"context"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type Storage interface {
	Client() ClientService
	Contractor() ContractorService
}

type ClientService interface {
	CreateTender(ctx context.Context, req *pb.CreateTenderReq) (*pb.ResponseMessage, error)
	GetMyTenders(ctx context.Context, req *pb.GetMyTendersReq) (*pb.TendersList, error)
	GetAllTenders(ctx context.Context, req *pb.GetAllTendersReq) (*pb.TendersList, error)
	UpdateTenderStatus(ctx context.Context, req *pb.UpdateTenderReq) (*pb.ResponseMessage, error)
	DeleteTender(ctx context.Context, req *pb.DeleteTenderReq) (*pb.ResponseMessage, error)
	SelectWinner(ctx context.Context, req *pb.SelectWinnerReq) (*pb.ResponseMessage, error)
}

type ContractorService interface {
	SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error)
	GetBidsForTender(ctx context.Context, req *pb.GetBidsRequest) (*pb.BidsListResponse, error)
}
