package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	Nama      string `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	Harga     int    `gorm:"type:int;not null" json:"harga" form:"harga"`
	Deskripsi string `gorm:"type:varchar(255);not null" json:"deskripsi" form:"deskripsi"`
	Longitude int    `gorm:"type:int;not null" json:"long" form:"long"`
	Latitude  int    `gorm:"type:int;not null" json:"lat" form:"lat"`
}

type Get_Rooms struct {
	gorm.Model
	Nama      string `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	Harga     int    `gorm:"type:int;not null" json:"harga" form:"harga"`
	Deskripsi string `gorm:"type:varchar(255);not null" json:"deskripsi" form:"deskripsi"`
	Longitude int    `gorm:"type:int;not null" json:"long" form:"long"`
	Latitude  int    `gorm:"type:int;not null" json:"lat" form:"lat"`
}
