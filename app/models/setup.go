package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database_string := "host=localhost user=postgres password=1234ASDFjkl; dbname=ProductivityApp port=5432"
	database, err := gorm.Open(postgres.Open(database_string), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Contact{})
	if err != nil {
		return
	}

	DB = database
}
