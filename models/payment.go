package models

import (
	"time"

	"gorm.io/gorm"
)

// Payment represents the Payments table in the database
type Payment struct {
	gorm.Model
	TransactionID uint      `json:"transaction_id" gorm:"unique;not null"`
	TglBayar      time.Time `json:"tgl_bayar" form:"tgl_bayar"`
	TotalBayar    float64   `json:"total_bayar" form:"total_bayar"`
}

// AfterSave callback to update Transaction Keterangan after payment is created
func (p *Payment) AfterSave(tx *gorm.DB) error {
	var transaction Transaction
	if err := tx.Model(&Transaction{}).Where("id = ?", p.TransactionID).First(&transaction).Error; err == nil {
		// Update the Transaction Keterangan to "Success"
		transaction.Keterangan = StatusSuccess
		tx.Save(&transaction)
	}
	return nil
}
