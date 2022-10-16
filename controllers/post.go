package controllers

import (
	"net/http"

	"github.com/kasfulk/golang-echo-mysql/models"
	"github.com/labstack/echo/v4"
)

func PostIndex(c echo.Context) error {
	DB := models.DB
	result, err := DB.QueryString("SELECT * FROM Post")
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err)
	}
	return c.JSON(http.StatusOK, result)
}
