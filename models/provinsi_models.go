package models

import "gorm.io/gorm"

type Provinsi struct {
	gorm.Model
	Nama_Provinsi string `gorm:"type:varchar(255);not null" json:"nama_provinsi" form:"nama_provinsi"`
	Kota          []Kota `gorm:"foreignKey:ProvinsiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
