package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Tanggal    time.Time `json:"tanggal" form:"tanggal"`
	Keterangan string    `json:"keterangan" form:"keterangan"`
	UserID     uint      `json:"user_id" form:"user_id"`
	Payment    Payment   `gorm:"foreignKey:TransactionID"`
}
