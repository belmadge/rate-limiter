package utils

import (
	"net/http"

	"github.com/belmadge/rate-limiter/infra/handler"
	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(rateLimiter *handler.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")

		if token != "" && !rateLimiter.AllowToken(c, token) {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
			return
		}

		if token == "" && !rateLimiter.AllowIP(c, ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
