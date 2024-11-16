package app

import (
	"fmt"
	"github.com/Hackaton-UDEVS/tender-service/pkg/kafka"

	"net"

	"github.com/Hackaton-UDEVS/tender-service/internal/config"
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
	db, err := postgres.ConnectionPostgres()
	if err != nil {
		log.Error("Error with conenct postgres", zap.Error(err))
		return
	}

	app := service.NewTenderService(db)

	kafka_handler := kafka.KafkaHandler{
		App: app,
	}

	err = kafka.Register(&kafka_handler)
	if err != nil {
		log.Error("Error with register kafka", zap.Error(err))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.FIRSTHOST, cfg.FIRSTPORT))
	if err != nil {
		log.Error("Error with listen", zap.Error(err))
		return
	}
	defer listener.Close()
	s := grpc.NewServer()

	if err := s.Serve(listener); err != nil {
		log.Error("Error while initializing server", zap.Error(err))
		return
	}
}
