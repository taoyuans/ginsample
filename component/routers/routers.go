package routers

import (
	"fmt"
	"ginsample/component/apis"
	"ginsample/config"
	"ginsample/lib/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(gormDB *gorm.DB) *gin.Engine {
	gin.SetMode(config.ConfigValue.GinMode)

	r := gin.New()

	r.Static("/static", "static")

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Logger())
	// r.Pre(middleware.RemoveTrailingSlash())
	r.Use(gin.Recovery())
	// r.Use(cors.New(auth.CorsConfig(config)))
	// e.Use(middleware.Logger())
	r.Use(middleware.SetRequestID())
	r.Use(middleware.SetDBMiddleware(gormDB))

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
