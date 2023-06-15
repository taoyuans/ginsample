package models

import (
	"context"

	"ginsample/lib/errs"
	"ginsample/lib/factory"
	"ginsample/lib/gohash"

	"gorm.io/gorm"
)

type User struct {
	Id       int64  `json:"id"  gorm:"primaryKey"`
	UserId   int64  `json:"userId" gorm:"unique"`
	PhoneNo  string `json:"phoneNo"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Deleted  bool   `json:"deleted"`
	BaseModel
	Version int `json:"version" gorm:"version"`
}

func (User) TableName() string {
	return "users"
}

func (User) GetByUserId(ctx context.Context, userId int64) (*User, error) {
	var user User
	err := factory.DB(ctx).Where("user_id = ?", userId).Where("deleted = 0").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errs.Trace(err)
	}
	return &user, nil
}

func (u *User) Create(ctx context.Context) (int64, error) {
	result := factory.DB(ctx).Create(u)
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

func (User) CheckPassword(ctx context.Context, userName, password string) (*User, error) {
	var user User
	err := factory.DB(ctx).Where("user_name = ? or phone_no = ? or email = ? ", userName, userName, userName).Where("password = ? ", gohash.Hash256(password)).Where("deleted = 0").First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errs.Trace(err)
	}
	return &user, nil
}
