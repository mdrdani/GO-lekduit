package routes

import (
	"GO-lekduit/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// route Users
	e.GET("/users", controllers.GetUserController)

	// route transactions
	e.GET("/transactions", controllers.GetTransactionController)

	return e
}
