package models

import (
	"time"

	"gorm.io/gorm"
)

// Payment represents the Payments table in the database
type Payment struct {
	gorm.Model
	TglBayar   time.Time `json:"tgl_bayar" form:"tgl_bayar"`
	TotalBayar float64   `json:"total_bayar" form:"total_bayar"`
}
