package ratelimiter

import (
	"context"
	"time"

	"github.com/ShervilG/spinboard-go/rediscache"
)

func IsRateLimited(key string, limit int, ttl time.Duration) bool {
	redisClient := rediscache.GetRedisClient()
	if redisClient == nil {
		return false
	}

	res := redisClient.SetNX(context.Background(), key, 1, ttl)
	if res.Val() {
		return false
	}

	intRes := redisClient.Incr(context.Background(), key)
	return intRes.Val() > int64(limit)
}
