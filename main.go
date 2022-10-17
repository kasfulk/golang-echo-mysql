package main

import (
	"github.com/joho/godotenv"
	"github.com/kasfulk/golang-echo-mysql/models"
	"github.com/kasfulk/golang-echo-mysql/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.ConnectDatabase()
	routes.BaseRoute()
}
