package routes

import (
	"GO-lekduit/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	// route login
	e.POST("/login", controllers.LoginController)

	AuthRoute := e.Group("")
	AuthRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	// route Users
	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.GetUserByIDController)
	e.POST("/users", controllers.AddUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	// route transactions
	AuthRoute.GET("/transactions", controllers.GetTransactionController)
	AuthRoute.GET("/transactions/:id", controllers.GetTransactionByIDController)
	AuthRoute.POST("/transactions", controllers.AddTransactionController)
	AuthRoute.PUT("/transactions/:id", controllers.UpdateTransactionController)

	// route payments
	AuthRoute.GET("/payments", controllers.GetPaymentController)
	AuthRoute.POST("/payments", controllers.AddPaymentsController)

	return e
}
