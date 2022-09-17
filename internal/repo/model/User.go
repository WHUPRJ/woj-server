package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	UserName   string `json:"user_name"  gorm:"not null;uniqueIndex"`
	NickName   string `json:"nick_name"  gorm:"not null"`
	Password   []byte `json:"-"          gorm:"not null"`
	IsEnabled  bool   `json:"is_enabled" gorm:"not null;index"`
}
