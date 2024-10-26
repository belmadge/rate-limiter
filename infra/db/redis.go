package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitializeRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	return rdb
}
