package controllers

import (
	"GO-lekduit/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetTransactionController(c echo.Context) error {
	transactions, e := database.GetTransactions()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       "success",
		"transactions": transactions,
	})
}
