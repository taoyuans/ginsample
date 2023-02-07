package routers

import (
	"ginsample/component/apis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouters(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ginsample is start work")
	})
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := g.Group("/api/v1")
	v1.GET("/users", apis.UserApi{}.GetUsers)
}
