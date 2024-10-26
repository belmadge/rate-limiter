package handler

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client RedisClient
}

type LimiterStore interface {
	Increment(ctx context.Context, key string) (int, error)
	SetExpiration(ctx context.Context, key string, expiration time.Duration) error
}

type RedisClient interface {
	Incr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
}

func NewRedisStore(client RedisClient) *RedisStore {
	return &RedisStore{client: client}
}

func (r *RedisStore) Increment(ctx context.Context, key string) (int, error) {
	count, err := r.client.Incr(ctx, key).Result()
	return int(count), err
}

func (r *RedisStore) SetExpiration(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}
