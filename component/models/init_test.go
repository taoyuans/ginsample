package models

import (
	"context"
	"flag"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ctx context.Context
var (
	appEnv = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
)

func init() {
	gormDB, err := gorm.Open(sqlite.Open("ginsample.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&User{})
	InitData(gormDB)

	ctx = context.WithValue(context.Background(), "DB", gormDB)

	defer RemoveSqlite()
}

func RemoveSqlite() {
	err := os.Remove("ginsample.db")
	if err != nil {
		fmt.Println(err)
	}
}
