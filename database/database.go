package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	// fmt.Println(os.Getenv("DB_ADDR"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		Username: os.Getenv("DB_USERNAME"),
		DB:       dbNo,
	})
	return rdb
}
