package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func InitMiddleware(r *gin.Engine) {
	logger, _ := zap.NewProduction()
	r.Use(Ginzap(logger, time.RFC3339, true))
	r.Use(RecoveryWithZap(logger, true))
	r.Use(CORS())
	r.Use(OperationLogger())
}
