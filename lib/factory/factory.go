package factory

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	// opentracing "github.com/opentracing/opentracing-go"
)

const ExportContextLogger = "ExportContextLogger"

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

func Logger(ctx context.Context) *logrus.Logger {
	v := ctx.Value("Logger")
	if v == nil {
		panic("Logger is not exist")
	}
	if logger, ok := v.(*logrus.Logger); ok {
		return logger
	}
	panic("Logger is not exist")
}

// func Redis(ctx context.Context) cache.Cache {
// 	v := ctx.Value("Redis")
// 	if v == nil {
// 		log.Println("redis is not exist")
// 	}
// 	if mycashe, ok := v.(cache.Cache); ok {
// 		return mycashe
// 	}
// 	panic("redis is not exist")
// }

// func Tracer(ctx context.Context) opentracing.Span {
// 	if s := opentracing.SpanFromContext(ctx); s != nil {
// 		return s
// 	}
// 	return opentracing.NoopTracer{}.StartSpan("")
// }
