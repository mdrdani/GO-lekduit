package config

import (
	"GO-lekduit/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic("error loading .env file")
	// }

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbName,
		dbPassword,
	)

	var err2 error
	DB, err2 = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Payment{}, &models.Transaction{})
}
