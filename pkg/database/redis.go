package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/okikechinonso/internals/ports"
	"os"
)

func NewRedisDB() ports.IRedis {
	return &datastore{}
}
func (d *datastore) ConnectRedisDB() *redis.Client {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "localhost:6379"
	}
	var redisClient = redis.NewClient(&redis.Options{
		Addr: url,
	})
	d.redis = redisClient
	return redisClient
}
