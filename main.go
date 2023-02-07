package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"ginsample/component/models"
	"ginsample/component/routers"
	configutil "ginsample/config"
	"ginsample/lib/errs"
	"ginsample/lib/goutils"
	"ginsample/lib/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mode   = flag.String("mode", "", "please input run mode.")
	appEnv = flag.String("env", "", "please input app env.")
)

func main() {
	flag.Parse()
	if *mode == "" || *appEnv == "" {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Printf(" - using mode:		mode = %s\n", *mode)
	fmt.Printf(" - using app_env:	app_env = %s\n", *appEnv)

	if !goutils.InArrayString(*appEnv, []string{"dev", "test", "prod"}) {
		fmt.Printf("[ERROR]  app_env=%s is not allowed.\n", *appEnv)
		os.Exit(1)
	}

	config := initConfigInformation()
	fmt.Printf("%+v", config)

	gormDB, err := gorm.Open(mysql.Open(config.DataBase.DataSourceName), &gorm.Config{})
	// gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormDB.AutoMigrate(&models.User{})

	switch *mode {
	case "api":
		if err := initGinApp(gormDB).Run(":" + config.HttpPort); err != nil {
			log.Println(errs.Trace(err))
			os.Exit(1)
		}
	case "init":
		models.InitData(gormDB)
	default:
		fmt.Printf("[ERROR]  mode=%s is not allowed.\n", *mode)
		os.Exit(1)
	}
}

func initGinApp(gormDB *gorm.DB) *gin.Engine {

	setSqlDBConfig(gormDB)
	r := InitRouter(gormDB)

	return r
}

func initConfigInformation() configutil.Config {
	var config configutil.Config
	if err := configutil.Read(*appEnv, &config); err != nil {
		panic(err)
	}

	return config
}

func setSqlDBConfig(gormDB *gorm.DB) {
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(20)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(100 * time.Second)
}

func InitRouter(gormDB *gorm.DB) *gin.Engine {
	gin.SetMode(configutil.ConfigValue.GinMode)

	r := gin.New()

	r.Static("/static", "static")

	r.Use(gin.Recovery())
	r.Use(middleware.SetRequestID())
	r.Use(middleware.SetDBMiddleware(gormDB))
	if configutil.ConfigValue.Log.Useable {
		r.Use(middleware.LogMiddleWare(configutil.ConfigValue.Log.LogFilePath, configutil.ConfigValue.Log.LogFileName))
	}

	routers.SetRouters(r)

	return r
}
