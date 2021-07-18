package cache

import (
	"basic-antd/init/config"
	"github.com/go-redis/redis/v8"
)

func RedisSetup() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
}
