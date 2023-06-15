package auth

import (
	"errors"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CheckToken(tokenString string) (jwt.MapClaims, error) {
	// 解析 JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 返回用于验证签名的密钥
		return []byte(JwtSecret), nil
	})

	// 检查解析和验证过程中是否发生错误
	if err != nil {
		fmt.Println("Check jwt token err :", err)
		return nil, err
	}

	// 检查 token 是否有效
	if !token.Valid {
		return nil, errors.New("invalid jwt token")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
