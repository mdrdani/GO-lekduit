package controllers

import (
	"GO-lekduit/lib/database"
	"GO-lekduit/models"
	"net/http"
	"strconv"

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
