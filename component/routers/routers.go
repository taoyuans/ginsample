package routers

import (
	"ginsample/component/apis"
	"ginsample/config"
	"ginsample/lib/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(gormDB *gorm.DB) *gin.Engine {
	gin.SetMode(config.ConfigValue.GinMode)

	r := gin.New()

	r.Static("/static", "static")

	r.Use(gin.Recovery())
	r.Use(middleware.SetRequestID())
	r.Use(middleware.SetDBMiddleware(gormDB))
	r.Use(middleware.LogMiddleWare())

	setRouters(r)

	return r
}

func setRouters(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ginsample is start work")
	})
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := g.Group("/api/v1")
	v1.GET("/users", apis.UserApi{}.GetUsers)
}
