package postgres

import (
	"database/sql"
	"fmt"

	logger "github.com/Hackaton-UDEVS/tender-service/internal/logger"

	pb "github.com/Hackaton-UDEVS/tender-service/internal/genproto/tender-service"
	"github.com/go-redis/redis"
)

type App1Repo struct {
	Rd *redis.Client
	Db *sql.DB
}

func NewApp1Repo(rd *redis.Client, db *sql.DB) *App1Repo {
	return &App1Repo{Rd: rd, Db: db}
}

func (t *App1Repo) TestCreate(req *pb.GetTendersReq) (*pb.GetTendersReq, error) {
	log, _ := logger.NewLogger()

	query := "insert into app(name) values ($1)"
	_, err := t.Db.Exec(query, "Avazbek")
	if err != nil {
		log.Error("Failed to insert into app")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully inserted new app")
	return &pb.GetTendersReq{}, nil
}
