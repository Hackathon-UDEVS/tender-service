package storage

import (
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
)

type Storage interface {
	App1() App1I
}

type App1I interface {
	TestCreate(req *pb.Test1Req) (*pb.Test1Res, error)
}
