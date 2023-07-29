package database

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
)

func GetPayments() (interface{}, error) {
	var payments []models.Payment

	if e := config.DB.Find(&payments).Error; e != nil {
		return nil, e
	}

	return payments, nil
}
