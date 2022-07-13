package apis

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ginsample/component/models"
	"ginsample/lib/middleware"
)

var (
	appEnv = flag.String("app-env", "test", "app env")
	r      = gin.New()
)

func init() {
	// var config configutil.Config
	// if err := configutil.Read(*appEnv, &config); err != nil {
	// 	panic(err)
	// }

	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&models.User{})
	models.InitData(gormDB)

	r.Static("/static", "static")
	r.Use(middleware.SetDBMiddleware(gormDB))
	InitApis(r)
	defer RemoveSqlite()
}

func RemoveSqlite() {
	err := os.Remove("ginsample.db")
	if err != nil {
		fmt.Println(err)
	}
}
