package models

type Kota struct {
	ID         int        `gorm:"primaryKey;autoIncrement;not null" json:"id" form:"id"`
	Nama_Kota  string     `gorm:"type:varchar(50);not null" json:"nama_kota" form:"nama_kota"`
	ProvinsiID uint       `json:"provinsi_id" form:"provinsi_id"`
	HomeStay   []HomeStay `gorm:"foreignKey:KotaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
