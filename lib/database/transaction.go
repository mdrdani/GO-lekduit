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

func GetTransactionByID(transactionID int) (models.Transaction, error) {
	var transaction models.Transaction

	if e := config.DB.First(&transaction, transactionID).Error; e != nil {
		return transaction, e
	}
	return transaction, nil
}
