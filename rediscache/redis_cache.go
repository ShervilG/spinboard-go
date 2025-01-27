package rediscache

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func SetupRedisClient() {
	redisUrl := os.Getenv("SPINBOARD_REDIS_URL")
	redisPass := os.Getenv("SPINBOARD_REDIS_PASS")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl + ":6379",
		Password: redisPass,
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		SetupRedisClient()
	}

	return redisClient
}
