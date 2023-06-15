package middleware

import (
	"context"

	"ginsample/lib/logger"

	"github.com/gin-gonic/gin"
)

func SetLogMiddleWare(fileName string, maxSize, maxBackups, maxAge int, logLevel int8) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		logContext := logger.New(req, fileName, maxSize, maxBackups, maxAge, logLevel)
		ctx.Request = req.WithContext(context.WithValue(req.Context(), "LogContext", logContext))
		// startTime := time.Now()

		//记录接口调用开始日志
		logContext.Info().
			Int("Status", ctx.Writer.Status()).
			Msg("request.")

		ctx.Next()

		// endTime := time.Now()
		// // 记录接口调用结束日志 + 计算执行时间
		// logContext.Info().
		// 	Dur("latencyTime", endTime.Sub(startTime)).
		// 	Int("Status", ctx.Writer.Status()).
		// 	Msg("request end")
	}
}
