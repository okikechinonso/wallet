package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient = ConnectRedisDB()

func ConnectRedisDB() *redis.Client {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "localhost:6379"
	}
	var redisClient = redis.NewClient(&redis.Options{
		Addr: url,
	})

	return redisClient
}
