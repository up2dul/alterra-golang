package route

import (
	"alterra2/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/books", controller.AddBook)
	e.GET("/books", controller.GetAllBooks)
	e.GET("/books/:id", controller.GetBook)
	return e
}
