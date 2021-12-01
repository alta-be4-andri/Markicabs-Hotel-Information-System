package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	RoomsID    uint   `json:"rooms_id" form:"rooms_id"`
	Nama_Photo string `gorm:"type:varchar(50);not null" json:"nama_photo" form:"nama_photo"`
	Url        string `gorm:"type:longtext" json:"url" form:"url"`
}

type Get_Photo struct {
	RoomsID    int
	Nama_Photo string
}
