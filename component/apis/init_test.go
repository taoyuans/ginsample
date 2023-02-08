package apis

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ginsample/component/models"
	configutil "ginsample/config"
	"ginsample/lib/middleware"
)

var (
	appEnv = flag.String("app-env", "test", "app env")
	r      *gin.Engine
)

func init() {
	var config configutil.Config
	if err := configutil.Read(*appEnv, &config); err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer RemoveSqlite()
	gormDB.AutoMigrate(&models.User{})
	models.InitData(gormDB)

	gin.SetMode(configutil.ConfigValue.GinMode)

	r = gin.New()

	r.Static("/static", "static")
	r.Use(middleware.SetDBMiddleware(gormDB))
	r.Use(middleware.SetLogMiddleWare("../../logs", "log.log"))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ginsample is start work")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")
	v1.GET("/users", UserApi{}.GetUsers)

}

func RemoveSqlite() {
	err := os.Remove("ginsample.db")
	if err != nil {
		fmt.Println(err)
	}
}
