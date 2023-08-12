package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"realName" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:256"`
	Mobile   string `json:"mobile" gorm:"size:64"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"size:128;not null"`
}
