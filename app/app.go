package app

import (
	"fmt"
	"net"

	"github.com/Hackaton-UDEVS/tender-service/internal/config"
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	logger "github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/service"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Run() {
	log, err := logger.NewLogger()

	cfg := config.Load()
	if err != nil {
		fmt.Println("Error with", err)
		return
	}

	// Storage
	stg, err := postgres.ConnectionPostgres()
	if err != nil {
		log.Error("Error with conenct postgres", zap.Error(err))
		return
	}

	// Connection
	lis, err := net.Listen("tcp", cfg.HTTP_PORT)
	if err != nil {
		fmt.Printf("Error while creating TCP listener: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	client_serivce := service.NewClientService(stg)
	contractor_service := service.NewContractorService(stg)

	// Registering services
	pb.RegisterClientServiceServer(server, client_serivce)
	pb.RegisterContractorServiceServer(server, contractor_service)

	// Run
	fmt.Println("Server listening at", cfg.HTTP_PORT)
	if server.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ")
	}
}
