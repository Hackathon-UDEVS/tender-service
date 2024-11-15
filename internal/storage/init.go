package storage

import (
	pb "github.com/Hackaton-UDEVS/first/internal/genproto/first-service"
)

type Storage interface {
	App1() App1I
}

type App1I interface {
	TestCreate(req *pb.Test1Req) (*pb.Test1Res, error)
}
