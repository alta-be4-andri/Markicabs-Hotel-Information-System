package models

import "gorm.io/gorm"

type Kota struct {
	gorm.Model
	NamaKota   string `gorm:"type:varchar(50);not null" json:"nama_provinsi" form:"nama_provinsi"`
	ProvinsiID uint   `json:"id_provinsi" form:"id_provinsi"`
}
