package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	var DBUser = os.Getenv("DB_USERNAME")
	var DBPass = os.Getenv("DB_PASSWORD")
	var DBHost = os.Getenv("DB_HOST")
	var DBPort = os.Getenv("DB_PORT")
	var DBName = os.Getenv("DB_NAME")

	ConnectionString := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(ConnectionString))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})
	DB = db
}
