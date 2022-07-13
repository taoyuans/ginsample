package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"-" gorm:"created"`
	UpdatedAt time.Time `json:"-" gorm:"updated"`
}
