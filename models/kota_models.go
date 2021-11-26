package models

type Kota struct {
	Id         int
	Nama_Kota  string `gorm:"type:varchar(50);not null" json:"nama_provinsi" form:"nama_provinsi"`
	ProvinsiID uint   `json:"id_provinsi" form:"id_provinsi"`
}
