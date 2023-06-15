package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// 建立一個 middleware 在所有的 response header 中加入 "X-Request-Id" 便于追溯
func SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get("TraceId")
		if len(requestId) <= 0 {
			uuidV4, err := uuid.NewV4()
			if err != nil {
				fmt.Printf("middleware tlogTraceId func err :%v", err)
			}
			requestId = uuidV4.String()
			ctx.Request.Header.Add("TraceId", requestId)
		}

		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), "TraceId", requestId))
		ctx.Next()
	}
}
