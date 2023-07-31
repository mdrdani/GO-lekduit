package controllers

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func createToken(user models.User) (string, error) {
	// Buat token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims JWT
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 1 hari

	// Menandatangani token dan mendapatkan string representasi dari token tersebut
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY"))) // Ganti "secret_key" dengan key rahasia Anda
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginController(c echo.Context) error {
	// Ambil data dari request
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Cari user berdasarkan email
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid email or password",
		})
	}

	// Verifikasi password
	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid email or password",
		})
	}

	// Buat token JWT
	token, err := createToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create JWT token",
		})
	}

	// Kirim token sebagai respons
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
