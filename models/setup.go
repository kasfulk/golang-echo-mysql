package models

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func ConnectDatabase() {
	var err error

	var DBUser = os.Getenv("DB_USERNAME")
	var DBPass = os.Getenv("DB_PASSWORD")
	var DBHost = os.Getenv("DB_HOST")
	var DBPort = os.Getenv("DB_PORT")
	var DBName = os.Getenv("DB_NAME")

	ConnectionString := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8"
	db, err := xorm.NewEngine("mysql", ConnectionString)
	if err != nil {
		panic(err)
	}
	DB = db
	db.Sync2(new(Post))
}
