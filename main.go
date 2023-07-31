package main

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
	"GO-lekduit/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	// Inisialisasi user pertama kali
	user := models.User{
		Nama:     "Administrator",
		Email:    "admin@golekduit.com",
		Alamat:   "Everywhere",
		NoTelpon: "123456789",
		Password: "secret123",
	}

	// Menyimpan user ke dalam database
	result := config.DB.Create(&user)
	if result.Error != nil {
		panic("Failed to create the first user")
	}

	e := echo.New()
	e = routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8000"))
}
