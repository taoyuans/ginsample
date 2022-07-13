package models

import (
	"fmt"

	"gorm.io/gorm"
)

func InitData(gorm *gorm.DB) {
	var (
		users = []User{
			{Id: 1, Code: "xiaoming", Name: "小明", Enable: true},
			{Id: 2, Code: "xiaozhang", Name: "小张", Enable: true},
		}
	)
	for _, u := range users {
		result := gorm.Create(&u)
		if result.Error != nil {
			fmt.Println("init user err: ", result.Error)
		}
	}
}
