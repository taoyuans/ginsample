package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetMysqlMiddleware(gormDB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), "DB", gormDB))
		ctx.Next()
	}
}

func SetRedisMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), "Redis", redis))
		ctx.Next()
	}
}
