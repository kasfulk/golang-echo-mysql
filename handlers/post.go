package handlers

import (
	"net/http"

	"github.com/kasfulk/golang-echo-mysql/databases/models"
	"github.com/labstack/echo/v4"
)

func PostIndex(c echo.Context) error {
	posts := models.ShowPost()
	return c.JSON(http.StatusOK, posts)
}

func PostDetail(c echo.Context) error {
	id := c.Param("id")
	post, err := models.ShowPostDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

// func PostCreate(c echo.Context) error {

// }

// func PostUpdate(c echo.Context) error {

// }

// func PostDelete(c echo.Context) error {

// }
