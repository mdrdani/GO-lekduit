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
	e.POST("/transactions", controllers.AddTransactionController)
	e.PUT("/transactions/:id", controllers.UpdateTransactionController)

	// route payments
	e.GET("/payments", controllers.GetPaymentController)
	e.POST("/payments", controllers.AddPaymentsController)

	return e
}
