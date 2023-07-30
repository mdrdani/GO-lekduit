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
	e.POST("/users", controllers.AddUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	// route transactions
	e.GET("/transactions", controllers.GetTransactionController)
	e.GET("/transactions/:id", controllers.GetTransactionByIDController)

	// route payments
	e.GET("/payments", controllers.GetPaymentController)

	return e
}
