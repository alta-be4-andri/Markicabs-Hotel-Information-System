package models

import "gorm.io/gorm"

type HomeStay struct {
	gorm.Model
	Nama    string `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	UsersID uint   `json:"user_id" form:"user_id"`
	KotaID  uint   `json:"kota_id" form:"kota_id"`
	Bujur   string `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Lintang string `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
	Alamat  string `gorm:"type:varchar(100);not null" json:"address" form:"lat"`
}

type Get_HomeStay struct {
	Nama      string
	Harga     int
	Deskripsi string
	Longitude int
	Latitude  int
}
