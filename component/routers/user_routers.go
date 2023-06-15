package routers

import (
	"ginsample/component/apis"

	"github.com/gin-gonic/gin"
)

func SetUserRouters(g *gin.Engine) {
	v1 := g.Group("/api/v1/users")
	v1.GET("/", apis.UserApi{}.GetByUserId)

	//账号密码登录
	v1.POST("/login-user-name", apis.UserApi{}.LoginByUserName)
}
