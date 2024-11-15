package app

import (
	"fmt"
	"github.com/Hackaton-UDEVS/first/internal/config"
	pb "github.com/Hackaton-UDEVS/first/internal/genproto/first-service"
	logger "github.com/Hackaton-UDEVS/first/internal/logger"
	"github.com/Hackaton-UDEVS/first/internal/service"
	"github.com/Hackaton-UDEVS/first/internal/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func Run() {
	log, err := logger.NewLogger()
	cfg := config.Load()
	if err != nil {
		fmt.Println("Error with", err)
		return
	}
	db, err := postgres.ConnectionPostgres()
	if err != nil {
		log.Error("Error with conenct postgres", zap.Error(err))
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.FIRSTHOST, cfg.FIRSTPORT))
	if err != nil {
		log.Error("Error with listen", zap.Error(err))
		return
	}
	defer listener.Close()
	s := grpc.NewServer()
	app1 := service.NewApp1Service(db)
	pb.RegisterFirstServiceServer(s, app1)
	log.Info(fmt.Sprintf("Server start on port: %d", cfg.FIRSTPORT))

	if err := s.Serve(listener); err != nil {
		log.Error("Error while initializing server", zap.Error(err))
		return
	}
}
