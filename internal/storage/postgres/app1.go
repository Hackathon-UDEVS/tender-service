package postgres

import (
	"database/sql"
	"fmt"
	pb "github.com/Hackaton-UDEVS/first/internal/genproto/first-service"
	"github.com/Hackaton-UDEVS/first/internal/logger"
	"github.com/go-redis/redis"
)

type App1Repo struct {
	Rd *redis.Client
	Db *sql.DB
}

func NewApp1Repo(rd *redis.Client, db *sql.DB) *App1Repo {
	return &App1Repo{Rd: rd, Db: db}
}

func (t *App1Repo) TestCreate(req *pb.Test1Req) (*pb.Test1Res, error) {
	log, _ := logger.NewLogger()

	query := "insert into app(name) values ($1)"
	_, err := t.Db.Exec(query, "Avazbek")
	if err != nil {
		log.Error("Failed to insert into app")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully inserted new app")
	return &pb.Test1Res{
		Meassage: req.Meassage,
	}, nil
}
