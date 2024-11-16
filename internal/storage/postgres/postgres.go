package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Hackaton-UDEVS/tender-service/internal/config"
	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	logger "github.com/Hackaton-UDEVS/tender-service/internal/logger"
	"github.com/Hackaton-UDEVS/tender-service/internal/storage"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Storage struct {
	Db   *sql.DB
	Rd   *redis.Client
	App1 storage.App1I
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
	app := NewApp1Repo(rd, db)
	app.TestCreate(&pb.Test1Req{Meassage: "hjsdg"})
	return &Storage{
		Db:   db,
		Rd:   rd,
		App1: app,
	}, nil
}

func (stg *Storage) User() *storage.App1I {
	if stg.App1 == nil {
		stg.App1 = NewApp1Repo(stg.Rd, stg.Db)
	}
	return &stg.App1
}
