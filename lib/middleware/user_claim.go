package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"ginsample/lib/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SkipperPaths = []string{
	"/",
	// "/ping",
	// "/health.json",
	"/api/v1/users/login-user-name",
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		var tokenString string
		if authHeader := ctx.Request.Header.Get("Authorization"); len(authHeader) > 0 && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = authHeader[7:]

			// 解析 JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// 返回用于验证签名的密钥
				return []byte(auth.JwtSecret), nil
			})

			// 检查解析和验证过程中是否发生错误
			if err != nil {
				ctx.Abort()
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.Error(),
				})
				return
			}

			// 检查 token 是否有效
			if !token.Valid {
				ctx.Abort()
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "jwt token is empty!",
				})

				return
			}
			claims := token.Claims.(jwt.MapClaims)
			req = req.WithContext(context.WithValue(req.Context(), "Claim", claims))
			req = req.WithContext(context.WithValue(req.Context(), "Token", token))
			ctx.Request = req

		} else {
			if !IsJWTMiddlewareSkipperPath(ctx.Request.URL.Path) {
				ctx.Abort()
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "jwt token is empty!",
				})
				return
			}
		}
		ctx.Next()
	}
}

func IsJWTMiddlewareSkipperPath(path string) bool {
	if len(path) > 0 {
		for _, skipperPath := range SkipperPaths {
			if path == skipperPath {
				return true
			}
		}
	}
	return false
}

func UserClaimMiddelware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		var token string
		if authHeader := ctx.Request.Header.Get("Authorization"); len(authHeader) > 0 && strings.HasPrefix(authHeader, "Bearer ") {
			token = authHeader[7:]
			si := strings.Index(token, ".")
			li := strings.LastIndex(token, ".")
			if si == -1 || li == -1 || si == li {
				ctx.Next()
			}

			payload := token[si+1 : li]
			if payload == "" {
				ctx.Next()
			}

			payloadBytes, err := decodeSegment(payload)
			if err != nil {
				ctx.Next()
			}

			var user map[string]interface{}

			err = json.Unmarshal(payloadBytes, &user)
			if err != nil {
				fmt.Println(err)
				ctx.Next()
			}
			req = req.WithContext(context.WithValue(req.Context(), "claims", user))
			// ctx.Request = req.WithContext(context.WithValue(req.Context(), "Token", token))
			req = req.WithContext(context.WithValue(req.Context(), "token", token))
			ctx.Request = req
		}

		ctx.Next()

	}
}

func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
