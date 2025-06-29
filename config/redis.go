package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func Redis() *redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(opt)

	return rdb
}
