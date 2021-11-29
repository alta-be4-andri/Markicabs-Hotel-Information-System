package models

import "gorm.io/gorm"

// struct HomeStay
type HomeStay struct {
	gorm.Model
	Nama      string   `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	UsersID   uint     `json:"users_id" form:"users_id"`
	KotaID    int      `json:"kota_id" form:"kota_id"`
	Longitude float64  `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Latitude  float64  `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
	Alamat    string   `gorm:"type:longtext;not null" json:"alamat" form:"alamat"`
	Rating    float64  `json:"rating" form:"rating"`
	Rooms     []Rooms  `gorm:"foreignKey:HomeStayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Review    []Review `gorm:"foreignKey:HomeStayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Get_HomeStay struct {
	Nama      string
	Longitude float64
	Latitude  float64
	Alamat    string
	Rating    float64
}
