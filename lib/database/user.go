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

	if e := config.DB.Preload("Transactions.Payment").First(&user, userID).Error; e != nil {
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

func UpdateUser(userID int, updatedUser models.User) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, userID).Error; e != nil {
		return user, e
	}

	user.Nama = updatedUser.Nama
	user.Email = updatedUser.Email
	user.Alamat = updatedUser.Alamat
	user.NoTelpon = updatedUser.NoTelpon

	if e := config.DB.Save(&user).Error; e != nil {
		return user, e
	}

	return user, nil
}

func DeleteUser(userID int) error {
	var user models.User
	if e := config.DB.First(&user, userID).Error; e != nil {
		return e
	}

	if e := config.DB.Delete(&user).Error; e != nil {
		return e
	}

	return nil
}
