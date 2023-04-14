package controller

import (
	"alterra2/config"
	"alterra2/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) error {
	var book model.Book
	c.Bind(&book)

	result := config.DB.Create(&book)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"success", book,
	})
}

func GetAllBooks(c echo.Context) error {
	var books []model.Book

	result := config.DB.Find(&books)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"success", books,
	})
}

func GetBook(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, model.Response{
		"success", id,
	})
}
