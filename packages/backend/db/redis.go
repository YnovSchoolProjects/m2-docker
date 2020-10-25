package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

func GetRedisConnection() (*redis.Client, context.Context) {
	dsn := os.Getenv("REDIS_DSN")

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb, ctx
}