package main

import (
	"fmt"

	"github.com/belmadge/rate-limiter/config"
	"github.com/belmadge/rate-limiter/infra/db"
	"github.com/belmadge/rate-limiter/infra/handler"
)

func main() {
	cfg := config.LoadConfig()
	rdb := db.InitializeRedis()

	store := handler.NewRedisStore(rdb)
	rateLimiter := handler.NewRateLimiter(store, cfg.RateLimitIP, cfg.RateLimitToken, cfg.BlockDuration)

	router := setupRouter(rateLimiter)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Server startup error:", err)
	}
}
