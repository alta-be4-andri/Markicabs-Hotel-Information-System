package models

import "gorm.io/gorm"

type Fasilitas struct {
	gorm.Model
	Nama_Fasilitas  string          `gorm:"type:varchar(255);not null" json:"nama_fasilitas" form:"nama_fasilitas"`
	FasilitasRoomID []FasilitasRoom `gorm:"foreignKey:FasilitasID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
