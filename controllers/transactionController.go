package controllers

import (
	"GO-lekduit/lib/database"
	"net/http"
	"strconv"

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

func GetTransactionByIDController(c echo.Context) error {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Transaction ID")
	}

	transaction, err := database.GetTransactionByID(transactionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Transaction Not Found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"transaction": transaction,
	})
}
