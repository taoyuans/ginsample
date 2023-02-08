package models

import (
	"context"
	"fmt"
	configutil "ginsample/config"
	"ginsample/lib/errs"
	"ginsample/lib/factory"

	"gorm.io/gorm"
)

type User struct {
	Id     int64  `json:"id"  gorm:"primaryKey"`
	Code   string `json:"code" gorm:"index"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
	BaseModel
	Version int `json:"version" gorm:"version"`
}

func (User) TableName() string {
	return "users"
}

func (User) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := factory.DB(ctx).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errs.Trace(err)
	}

	//use config value
	fmt.Printf("【models】connfig value test service => %s\n", configutil.ConfigValue.Service)

	//add log
	factory.Logger(ctx).Info("【models】log test => log_info")

	return users, nil
}

func (a *User) Create(ctx context.Context) (int64, error) {
	result := factory.DB(ctx).Create(a)
	if result.Error != nil {
		return 0, errs.Trace(result.Error)
	}
	return result.RowsAffected, nil
}

func (User) GetById(ctx context.Context, id int64) (*User, error) {
	var user User
	err := factory.DB(ctx).Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errs.Trace(err)
	}
	return &user, nil
}
