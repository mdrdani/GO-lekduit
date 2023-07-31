package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string        `json:"nama" form:"nama"`
	Email        string        `json:"email" form:"email"`
	Alamat       string        `json:"alamat" form:"alamat"`
	NoTelpon     string        `json:"no_telpon" form:"no_telpon"`
	Password     string        `json:"password" form:"password"`
	Transactions []Transaction `gorm:"foreignkey:UserID"`
}
