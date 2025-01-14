package cache

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

var cache *ttlcache.Cache[string, string]

func SetupCache() {
	cache = ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](30 * time.Minute),
	)

	go cache.Start()
}

func Set(key string, value string, duration time.Duration) {
	cache.Set(key, value, duration)
}

func Get(key string) string {
	ret := cache.Get(key)
	if ret == nil {
		return ""
	}

	if ret.IsExpired() {
		cache.Delete(key)
		return ""
	}

	return ret.Value()
}
