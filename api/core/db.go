package models

import (
	"fmt"
	"ginapp/core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDatabase() *gorm.DB {
	dsn := getDBUrl()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to db")
		panic("Failed to connect database")
	} else {
		fmt.Println("Connect db success")

		database.AutoMigrate(&models.User{})
	}

	return database
}

func getDBUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_CHARSET"),
	)
}
