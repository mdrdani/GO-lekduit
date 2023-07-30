package controllers

import (
	"GO-lekduit/lib/database"
	"GO-lekduit/models"
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

func AddPaymentsController(c echo.Context) error {
	var newPayment models.Payment
	if err := c.Bind(&newPayment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Data")
	}

	payment, err := database.AddPayments(newPayment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add payment")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"payment": payment,
	})
}
