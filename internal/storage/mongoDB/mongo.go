package mongoDB

import (
	"context"
	logger "github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func ConnectionMongoDb() (*mongo.Client, error) {
	// Initialize logger
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	defer log.Sync()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB", zap.Error(err))
		return nil, err
	}

	log.Info("Connected to MongoDB successfully")
	return client, nil
}
