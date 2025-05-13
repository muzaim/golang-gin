package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	connectionStr := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=true&loc=Asia%%2FJakarta",
		ENV.DB_USERNAME,
		ENV.DB_PASSWORD,
		ENV.DB_URL,
		ENV.DB_DATABASE,
	)

	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB = db
}
