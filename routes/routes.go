package routes

import (
	"GO-lekduit/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)

	return e
}
