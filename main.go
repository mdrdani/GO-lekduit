package main

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
	"GO-lekduit/routes"
	"fmt"

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

	// Cari pengguna dengan email yang sama di dalam tabel User
	var existingUser models.User
	result := config.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		// Jika tidak ada pengguna dengan email yang sama, maka simpan pengguna ke dalam tabel
		result := config.DB.Create(&user)
		if result.Error != nil {
			panic("Failed to create the first user")
		}
	} else {
		// Jika sudah ada pengguna dengan email yang sama, tidak perlu menyimpan data pengguna baru
		// Anda bisa memberikan respons bahwa pengguna sudah ada, atau melakukan tindakan lain yang sesuai
		fmt.Println("User with the same email already exists.")
	}

	e := echo.New()
	e = routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8000"))
}
