package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"im/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	// 自动迁移
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	user := &models.User{}
	user.Name = "Fuck"
	// 创建一条项目
	db.Create(user)

	// read
	fmt.Println(db.First(user, 1))

	db.Model(user).Update("DeviceID", "test_device_id")
	fmt.Println("Test pass")

}
