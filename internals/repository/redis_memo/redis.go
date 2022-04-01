package redismemo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/okikechinonso/internals/ports"
	"strconv"
)

type redisRepository struct {
	Redis *redis.Client
}

func NewRedisRepository(client *redis.Client) ports.RedisRepository {
	return &redisRepository{
		Redis: client,
	}
}

func (r *redisRepository) Get(key string) (*float64, error) {
	var ctx = context.Background()
	val, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (r *redisRepository) Set(key string, data float64) error {
	var ctx = context.Background()
	err := r.Redis.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) Pub(str *string) {
	var ctx = context.Background()
	data := &struct {
		Balance string `json:"balance"`
	}{
		Balance: *str,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if err = r.Redis.Publish(ctx, "send-user-data", payload).Err(); err != nil {
		panic(err)
	}
}

func (r *redisRepository) Sub() {
	var ctx = context.Background()
	subscriber := r.Redis.Subscribe(ctx, "send-user-data")

	data := &struct {
		Balance string `json:"balance"`
	}{}
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), data); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")

		err = r.Redis.Set(ctx, "balance", data.Balance, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
