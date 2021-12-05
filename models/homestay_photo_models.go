package models

import (
	"gorm.io/gorm"
)

type HomeStayPhoto struct {
	gorm.Model
	HomeStayID uint   `json:"homestay_id" form:"homestay_id"`
	Nama_Photo string `gorm:"type:varchar(50);not null" json:"nama_photo" form:"nama_photo"`
	Url        string `gorm:"type:longtext" json:"url" form:"url"`
}

type Get_HomestayPhoto struct {
	HomeStayID int
	Nama_Photo string
}
