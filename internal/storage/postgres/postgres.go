package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Hackaton-UDEVS/tender-service/internal/config"
	"github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db         *sql.DB
	Rd         *redis.Client
	client     storage.ClientService
	contractor storage.ContractorService
}

// ConnectionPostgres initializes and returns a new Storage instance.
func ConnectionPostgres() (*Storage, error) {
	cfg := config.Load()

	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	// Create PostgreSQL connection string
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword,
		cfg.PostgresHost, cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}
	logs.Info("Successfully connected to PostgreSQL")

	// Initialize Redis client
	rd := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.REDIS_HOST, cfg.REDIS_PORT),
	})

	// Verify Redis connection
	_, err = rd.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	logs.Info("Successfully connected to Redis")

	return &Storage{
		Db: db,
		Rd: rd,
	}, nil
}

// Client returns a ClientService implementation.
func (s *Storage) Client() storage.ClientService {
	if s.client == nil {
		s.client = &ClientStorage{db: s.Db}
	}
	return s.client
}

// Contractor returns a ContractorService implementation.
func (s *Storage) Contractor() storage.ContractorService {
	if s.contractor == nil {
		s.contractor = &ContractorStorage{db: s.Db}
	}
	return s.contractor
}
