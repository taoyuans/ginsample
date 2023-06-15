package factory

import (
	"context"

	"ginsample/lib/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DB(ctx context.Context) *gorm.DB {
	v := ctx.Value("DB")
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*gorm.DB); ok {
		return db
	}
	panic("DB is not exist")
}

func Redis(ctx context.Context) *redis.Client {
	v := ctx.Value("Redis")
	if v == nil {
		panic("Redis is not exist")
	}
	if redis, ok := v.(*redis.Client); ok {
		return redis
	}
	panic("Redis is not exist")
}

func Logger(ctx context.Context) *logger.LogContext {
	v := ctx.Value("LogContext")
	if v == nil {
		panic("Logger is not exist")
	}
	if logger, ok := v.(*logger.LogContext); ok {
		return logger
	}
	panic("Logger is not exist")
}

func Token(ctx context.Context) string {
	v := ctx.Value("Token")
	if v == nil {
		return ""
	}
	if token, ok := v.(string); ok {
		return token
	}
	return ""
}
func Claim(ctx context.Context) map[string]interface{} {
	v := ctx.Value("Claim")
	if v == nil {
		return nil
	}
	if claim, ok := v.(map[string]interface{}); ok {
		return claim
	}
	return nil
}
func RequestId(ctx context.Context) string {
	v := ctx.Value("TraceId")
	if v == nil {
		panic("TraceId is not exist")
	}
	if traceId, ok := v.(string); ok {
		return traceId
	}
	panic("TraceId is not exist")
}
