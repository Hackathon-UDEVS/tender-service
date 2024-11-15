package kafka

import (
	"fmt"
	pb "github.com/Hackaton-UDEVS/first/internal/genproto/first-service"
	logger "github.com/Hackaton-UDEVS/first/internal/logger"
	"github.com/Hackaton-UDEVS/first/internal/service"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	_ "google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	App *service.App1Service
}

func (h *KafkaHandler) NewAppCreate() func(message []byte) {
	log, _ := logger.NewLogger()
	return func(message []byte) {
		var req pb.Test1Req
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Error("Cannot unmarshal JSON: %v", zap.Error(err))
			return
		}
		res, err := h.App.TestCreate(&req)
		if err != nil {
			log.Error("Cannot create app via Kafka: %v", zap.Error(err))
			return
		}
		log.Info(fmt.Sprintf("Created app: %+v", res))
	}
}
