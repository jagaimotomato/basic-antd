package redis

import (
	"basic-antd/pkg/redis"
	"context"
)

var (
	Ctx    = context.Background()
	Rdb, _ = redis.NewRedisClient()
)
