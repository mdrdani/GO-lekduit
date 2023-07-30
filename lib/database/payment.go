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

func AddPayments(newPayment models.Payment) (models.Payment, error) {
	if e := config.DB.Create(&newPayment).Error; e != nil {
		return newPayment, e
	}
	return newPayment, nil
}
