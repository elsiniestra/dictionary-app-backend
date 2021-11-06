package cache

import (
	"time"

	"github.com/SporkHubr/echo-http-cache"
	"github.com/SporkHubr/echo-http-cache/adapter/redis"
)

func New(redisURL string) (*cache.Client, error) {
	ringOpt := &redis.RingOptions{
		Addrs: map[string]string{
			"server": redisURL,
		},
	}
	return cache.NewClient(
		cache.ClientWithAdapter(redis.NewAdapter(ringOpt)),
		cache.ClientWithTTL(10*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
}
