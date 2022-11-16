package model

import "gorm.io/gorm"

type UserSql struct {
	gorm.Model
	User
}

type User struct {
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}
