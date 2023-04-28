package models

import (
	"fmt"
	"gorm.io/gorm"
	"im/utils"
)

type User struct {
	gorm.Model
	Name       string
	Password   string
	Phone      string
	Email      string
	Identity   string
	ClientIP   string
	ClientPort string
	LastLogin  uint64
	Heartbeat  uint64
	LastLogout uint64
	DeviceID   string
	sup        bool
}

func (u *User) TableName() string {
	return "users"
}

func GetUsers() []*User {
	data := make([]*User, 10)
	utils.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
