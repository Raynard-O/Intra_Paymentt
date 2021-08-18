package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.DB
	//ID int64 `json:"id"`
	Username string `json:"username"`
}