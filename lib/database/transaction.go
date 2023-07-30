package database

import (
	"GO-lekduit/config"
	"GO-lekduit/models"
)

func GetTransactions() (interface{}, error) {
	var transactions []models.Transaction

	if e := config.DB.Preload("Payment").Find(&transactions).Error; e != nil {
		return nil, e
	}

	return transactions, nil
}

func GetTransactionByID(transactionID int) (models.Transaction, error) {
	var transaction models.Transaction

	if e := config.DB.Preload("Payment").First(&transaction, transactionID).Error; e != nil {
		return transaction, e
	}
	return transaction, nil
}

func AddTransaction(newTransaction models.Transaction) (models.Transaction, error) {
	if e := config.DB.Create(&newTransaction).Error; e != nil {
		return newTransaction, e
	}
	return newTransaction, nil
}

func UpdateTransaction(transactionID int, updatedTransaction models.Transaction) (models.Transaction, error) {
	var transaction models.Transaction

	if e := config.DB.First(&transaction, transactionID).Error; e != nil {
		return transaction, e
	}

	transaction.Tanggal = updatedTransaction.Tanggal
	transaction.Keterangan = updatedTransaction.Keterangan
	transaction.UserID = updatedTransaction.UserID

	if e := config.DB.Save(&transaction).Error; e != nil {
		return transaction, e
	}

	return transaction, nil
}
