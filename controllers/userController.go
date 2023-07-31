package controllers

import (
	"GO-lekduit/lib/database"
	"GO-lekduit/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error)
	}

	// Create a custom JSON response structure
	type UserResponse struct {
		ID               uint   `json:"id"`
		Nama             string `json:"nama"`
		Email            string `json:"email"`
		Alamat           string `json:"alamat"`
		NoTelpon         string `json:"no_telpon"`
		TransactionCount int    `json:"transaction_count"`
	}

	// Create a slice to hold the custom JSON response
	var usersResponse []UserResponse

	// Loop through each user to calculate the transaction count
	for _, user := range users {
		transactionCount := len(user.Transactions)
		// Append the custom JSON response to the slice
		usersResponse = append(usersResponse, UserResponse{
			ID:               user.ID,
			Nama:             user.Nama,
			Email:            user.Email,
			Alamat:           user.Alamat,
			NoTelpon:         user.NoTelpon,
			TransactionCount: transactionCount,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  usersResponse,
	})
}

func GetUserByIDController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	user, err := database.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func AddUserController(c echo.Context) error {
	var newUser models.User
	if err := c.Bind(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid data")
	}

	user, err := database.AddUser(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add user")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func UpdateUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Data")
	}

	user, err := database.UpdateUser(userID, updatedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Update User")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success Data Update",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}

	err = database.DeleteUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to Delete User")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user deleted successfully",
	})
}
