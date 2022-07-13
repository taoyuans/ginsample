package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApis(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ginsample is start work")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")

	UserApi{}.Init(v1)
}

func (c UserApi) Init(g *gin.RouterGroup) {
	g.GET("/users", c.GetUsers)
}
