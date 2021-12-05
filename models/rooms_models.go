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
	RoomPhoto      []RoomPhoto     `gorm:"foreignKey:RoomsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BodyRoom struct {
	Nama_Room      string `gorm:"type:varchar(255);not null" json:"nama_room" form:"nama_room"`
	HomeStayID     uint   `json:"homestay_id" form:"homestay_id"`
	Total_Penghuni int    `gorm:"not null" json:"total_penghuni" form:"total_penghuni"`
	Harga          int    `gorm:"not null" json:"harga" form:"harga"`
	Deskripsi      string `gorm:"type:longtext;not null" json:"deskripsi" form:"deskripsi"`
	Fasilitas      []int  `json:"fasilitas" form:"fasilitas"`
}

type AllRoomsResponse struct {
	ID             uint
	Nama_Room      string
	HomeStayID     uint
	Total_Penghuni int
	Harga          int
	Deskripsi      string
	Url            string
}

type FasilitasRoomResponse struct {
	ID             uint
	Nama_Room      string
	HomeStayID     uint
	Total_Penghuni int
	Harga          int
	Deskripsi      string
	Url            string
	Fasilitas      []string
}
