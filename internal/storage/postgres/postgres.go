package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Hackaton-UDEVS/tender-service/internal/config"
	logger "github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Storage struct {
	Db     *sql.DB
	Rd     *redis.Client
	Tender storage.TenderServiceI
}

func ConnectionPostgres() (*Storage, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conf := config.Load()
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHOST, conf.DBPORT, conf.DBUSER, conf.DBPASSWORD, conf.DBNAME)
	fmt.Println(dns)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		logs.Error("Error while connecting postgres")
	}
	err = db.Ping()
	if err != nil {
		logs.Error("Error while pinging postgres", zap.Error(err))
	}
	logs.Info("Successfully connected to postgres")

	rd := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.REDISHOST, conf.REDISPORT),
	})
	tender := NewTenderRepo(db, rd)

	return &Storage{
		Db:     db,
		Rd:     rd,
		Tender: tender,
	}, nil
}

func (stg *Storage) User() *storage.TenderServiceI {
	if stg.Tender == nil {
		stg.Tender = NewTenderRepo(stg.Db, stg.Rd)
	}
	return &stg.Tender
}
