package controllers

import (
	"GO-lekduit/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetPaymentController(c echo.Context) error {
	payments, e := database.GetPayments()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"payments": payments,
	})
}
