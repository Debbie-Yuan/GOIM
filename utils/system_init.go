package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	// loading config file from config/app.yaml
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Can not open config file, please check it.")
	}

	//
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config database: ", viper.Get("database"))

}

func InitDatabase() {
	DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}
