package apis

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ginsample/component/models"
	cconfig "ginsample/config"
	"ginsample/lib/middleware"
)

var (
	appEnv = flag.String("app-env", "test", "app env")
	r      *gin.Engine
)

func init() {
	cconfig.ConfigValue = &cconfig.Config{}

	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&models.User{})
	models.InitData(gormDB)

	gin.SetMode(gin.DebugMode)

	r = gin.New()

	r.Static("/static", "static")
	r.Use(middleware.SetMysqlMiddleware(gormDB))
	r.Use(middleware.SetLogMiddleWare("./test.log", 1, 1, 1, 7))

	v1 := r.Group("/api/v1")
	v1.GET("/users", UserApi{}.GetByUserId)
	v1.POST("/users/login-user-name", UserApi{}.LoginByUserName)

}

func RemoveSqlite() {
	err := os.Remove("ginsample.db")
	if err != nil {
		fmt.Println(err)
	}
}
