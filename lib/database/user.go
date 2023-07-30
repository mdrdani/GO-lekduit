package database

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUserByID(userID int) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, userID).Error; e != nil {
		return user, e
	}
	return user, nil
}

func AddUser(newUser models.User) (models.User, error) {
	if e := config.DB.Create(&newUser).Error; e != nil {
		return newUser, e
	}
	return newUser, nil
}
