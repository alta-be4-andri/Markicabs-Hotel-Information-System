package models

type Provinsi struct {
	Id            int
	Nama_Provinsi string `gorm:"type:varchar(255);not null" json:"nama_provinsi" form:"nama_provinsi"`
	Kota          []Kota `gorm:"foreignKey:ProvinsiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
