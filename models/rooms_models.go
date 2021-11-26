package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	Nama_Room      string          `gorm:"type:varchar(255);not null" json:"nama_room" form:"nama_room"`
	HomeStayID     uint            `json:"homestay_id" form:"homestay_id"`
	Total_Penghuni int             `gorm:"not null" json:"total_penghuni" form:"total_penghuni"`
	Harga          int             `gorm:"not null" json:"harga" form:"harga"`
	Deskripsi      string          `gorm:"type:longtext;not null" json:"deskripsi" form:"deskripsi"`
	FasilitasRoom  []FasilitasRoom `gorm:"foreignKey:RoomsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
