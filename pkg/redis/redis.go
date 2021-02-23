package redis

import (
	"basic-antd/init/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func NewRedisClient() (*redis.Client, error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis failed")
		return nil, err
	}
	fmt.Printf("redis ping result: %s\n", pong)
	return RedisClient, nil
}
