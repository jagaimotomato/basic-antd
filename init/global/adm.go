package global

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Lg       *zap.Logger
	Eloquent *gorm.DB
	Version  = "1.0.0"
	Engine   = gin.New()
	// cache redis
	Rdb            *redis.Client
	Ctx            = context.Background()
	CasbinEnforcer *casbin.SyncedEnforcer
)
