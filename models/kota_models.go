package models

type Kota struct {
	ID         int        `gorm:"auto_increment;not null" json:"id" form:"id"`
	NamaKota   string     `gorm:"type:varchar(50);not null" json:"nama_kota" form:"nama_kota"`
	ProvinsiID int        `json:"id_provinsi" form:"id_provinsi"`
	HomeStay   []HomeStay `gorm:"foreignKey:KotaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
