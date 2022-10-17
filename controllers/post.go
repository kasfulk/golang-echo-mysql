package controllers

import (
	"net/http"

	"github.com/kasfulk/golang-echo-mysql/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PostIndex(c echo.Context) error {
	var posts []models.Post
	models.DB.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func PostDetail(c echo.Context) error {
	id := c.Param("id")
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "data tidak ditemkan!",
			})
		}
	}
	return c.JSON(http.StatusOK, post)
}
