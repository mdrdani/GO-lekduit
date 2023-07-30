package routes

import (
	"GO-lekduit/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// route Users
	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.GetUserByIDController)

	// route transactions
	e.GET("/transactions", controllers.GetTransactionController)

	// route payments
	e.GET("/payments", controllers.GetPaymentController)

	return e
}
