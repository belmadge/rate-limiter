package main

import (
	"github.com/belmadge/rate-limiter/infra/handler"
	"github.com/belmadge/rate-limiter/utils"
	"github.com/gin-gonic/gin"
)

func setupRouter(rateLimiter *handler.RateLimiter) *gin.Engine {
	router := gin.Default()
	router.Use(utils.RateLimiterMiddleware(rateLimiter))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "it is working"})
	})
	return router
}
