package controllers

import (
	"GO-lekduit/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
