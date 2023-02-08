package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// 建立一個 middleware 在所有的 response header 中加入 "X-Request-Id" 的項目
func SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuidV4, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("middleware SetRequestID func err :%v", err)
		}
		ctx.Request.Header.Add("X-Request-Id", uuidV4.String())

		ctx.Next()
	}
}
