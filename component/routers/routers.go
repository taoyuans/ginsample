package routers

import (
	"net/http"

	cconfig "ginsample/config"

	"github.com/gin-gonic/gin"
)

func SetBaseRouters(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ginsample is start work")
	})
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	g.GET("/config", func(c *gin.Context) {
		configValue := *cconfig.ConfigValue
		configValue.Mysql.DataSourceName = ""
		configValue.Redis.DataSourceName = ""
		configValue.Redis.Password = ""
		c.JSON(http.StatusOK, configValue)
	})
	g.GET("/health.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "UP"})
	})
	//prometheus 监控
	// g.GET("/actuator/prometheus", PromHandler(promhttp.Handler()))
}

// 添加prometheus Handler
func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
