package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kasfulk/golang-echo-mysql/models"
	// "github.com/kasfulk/golang-echo-mysql/models"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.ConnectDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
