package controllers

import (
	"GO-lekduit/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Id    int    `json:"Id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func LoginController(c echo.Context) error {
	var loginRequest models.LoginRequest
	c.Bind(&loginRequest)

	// Set custom claims
	claims := &jwtCustomClaims{
		"Muhamad Dani Ramanda",
		1,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success Token Accept ",
		"token":  t,
	})
}
