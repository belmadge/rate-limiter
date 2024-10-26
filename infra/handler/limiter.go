package handler

import (
	"context"
	"time"
)

type RateLimiter struct {
	store          LimiterStore
	rateLimitIP    int
	rateLimitToken int
	blockDuration  time.Duration
}

func NewRateLimiter(store LimiterStore, rateLimitIP, rateLimitToken, blockDuration int) *RateLimiter {
	return &RateLimiter{
		store:          store,
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockDuration:  time.Duration(blockDuration) * time.Second,
	}
}

func (rl *RateLimiter) allow(ctx context.Context, key string, limit int) bool {
	currentCount, err := rl.store.Increment(ctx, key)
	if err != nil {
		return false
	}

	if currentCount == 1 {
		err := rl.store.SetExpiration(ctx, key, rl.blockDuration)
		if err != nil {
			return false
		}
	}

	return currentCount <= limit
}

func (rl *RateLimiter) AllowToken(ctx context.Context, token string) bool {
	return rl.allow(ctx, "rate_limit_token:"+token, rl.rateLimitToken)
}

func (rl *RateLimiter) AllowIP(ctx context.Context, ip string) bool {
	return rl.allow(ctx, "rate_limit_ip:"+ip, rl.rateLimitIP)
}
