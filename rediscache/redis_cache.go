package rediscache

import (
	"context"
	"fmt"
	"os"

	"github.com/ShervilG/spinboard-go/redismessage"
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

	pubsub := redisClient.PSubscribe(context.Background(), "__keyevent@0__:expired")
	go func() {
		for {
			message, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				fmt.Printf("Error while receiving message from keyspace pubsub %v\n", err.Error())
			} else {
				redismessage.HandleRedisMessage(message)
			}
		}
	}()
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		SetupRedisClient()
	}

	return redisClient
}
