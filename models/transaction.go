package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionStatus string

const (
	StatusWaiting TransactionStatus = "Waiting"
	StatusSuccess TransactionStatus = "Success"
	StatusCancel  TransactionStatus = "Cancel"
)

type Transaction struct {
	gorm.Model
	Tanggal    time.Time         `json:"tanggal" form:"tanggal"`
	Keterangan TransactionStatus `json:"keterangan" form:"keterangan" gorm:"default:Waiting"`
	UserID     uint              `json:"user_id" form:"user_id"`
	Payment    Payment           `gorm:"foreignKey:TransactionID"`
}
