package database

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
)

func GetTransactions() (interface{}, error) {
	var transactions []models.Transaction

	if e := config.DB.Find(&transactions).Error; e != nil {
		return nil, e
	}

	return transactions, nil
}
