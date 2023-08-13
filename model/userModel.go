package model

import (
	"github.com/ikun666/go_webserver/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"realName" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:256"`
	Mobile   string `json:"mobile" gorm:"size:64"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"size:128;not null"`
}

// gorm hook func(*gorm.DB) error
func (u *User) BeforeCreate(tx *gorm.DB) error {
	password, err := utils.Crypt(u.Password)
	if err == nil {
		u.Password = password
	}
	return err
}
