package cache

import (
	"basic-antd/init/config"
	"basic-antd/init/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func setupRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	pong, err := global.Rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis failed")
	}
	fmt.Printf("redis ping result: %s\n", pong)
}
