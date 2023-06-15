package models

import (
	"fmt"

	"gorm.io/gorm"
)

func InitData(gorm *gorm.DB) {
	var (
		users = []User{
			{
				Id:       1,
				UserId:   10000000000000001,
				UserName: "zhangsan",
				Password: "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", //1234
				PhoneNo:  "13456789091",
				Email:    "13456789091@qq.com",
				Deleted:  false,
			},
			{
				Id:       2,
				UserId:   10000000000000002,
				UserName: "lisi",
				Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", //123456
				PhoneNo:  "17756789292",
				Email:    "17756789292@qq.com",
				Deleted:  false,
			},
		}
	)
	for _, u := range users {
		result := gorm.Create(&u)
		if result.Error != nil {
			fmt.Println("init user err: ", result.Error)
		}
	}
}
