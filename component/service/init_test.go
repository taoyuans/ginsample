package service

import (
	"context"
	"flag"
	"fmt"
	"os"

	"ginsample/component/models"
	cconfig "ginsample/config"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ctx context.Context
var (
	appEnv = flag.String("app-env", "test", "app env")
)

func init() {
	cconfig.ConfigValue = &cconfig.Config{}
	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&models.User{})
	models.InitData(gormDB)

	ctx = context.WithValue(context.Background(), "DB", gormDB)

}

func RemoveSqlite() {
	err := os.Remove("ginsample.db")
	if err != nil {
		fmt.Println(err)
	}
}
