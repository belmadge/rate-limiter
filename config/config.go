package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RateLimitIP    int
	RateLimitToken int
	BlockDuration  int
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	rateLimitIP := parseEnvAsInt("RATE_LIMIT_IP", 5)
	rateLimitToken := parseEnvAsInt("RATE_LIMIT_TOKEN", 10)
	blockDuration := parseEnvAsInt("BLOCK_DURATION", 300)

	return Config{
		RateLimitIP:    rateLimitIP,
		RateLimitToken: rateLimitToken,
		BlockDuration:  blockDuration,
	}
}

func parseEnvAsInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Invalid value for %s, defaulting to %d", key, defaultValue)
		return defaultValue
	}
	return value
}
