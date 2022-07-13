package main

import (
	"flag"
	"ginsample/component/apis"
	"ginsample/component/models"
	"ginsample/lib/errs"
	"ginsample/lib/filters"
	"log"
	"os"
	"sort"
	"time"

	configutil "ginsample/config"
	"ginsample/lib/middleware"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	appEnv = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
)

func main() {

	config := initConfigInformation()

	// gormDB, err := gorm.Open(mysql.Open(config.DataBase.DataSourceName), &gorm.Config{})
	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&models.User{})

	app := cli.NewApp()
	app.Name = "ginsample-api"
	app.Commands = []*cli.Command{
		{
			Name:    "api-server",
			Aliases: []string{"a"},
			Usage:   "run api server",
			Action: func(cliContext *cli.Context) error {
				if err := initGinApp(gormDB, cliContext).Run(":" + config.HttpPort); err != nil {
					log.Println(errs.Trace(err))
					return err
				}
				return nil
			},
		},
		{
			Name:    "init-data",
			Aliases: []string{"i"},
			Usage:   "init data",
			Action: func(c *cli.Context) error {
				models.InitData(gormDB)
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}

func initGinApp(gormDB *gorm.DB, cliContext *cli.Context) *gin.Engine {
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

	r := gin.New()

	r.Static("/static", "static")
	r.Use(gin.Logger())
	// r.Pre(middleware.RemoveTrailingSlash())
	r.Use(gin.Recovery())
	// r.Use(cors.New(auth.CorsConfig(config)))
	// e.Use(middleware.Logger())
	r.Use(middleware.SetRequestID())
	r.Use(filters.SetDBMiddleware(gormDB))

	apis.InitApis(r)

	return r
}

func initConfigInformation() configutil.Config {
	var config configutil.Config
	if err := configutil.Read(*appEnv, &config); err != nil {
		panic(err)
	}

	return config
}
