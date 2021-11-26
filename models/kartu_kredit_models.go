package models

import "gorm.io/gorm"

type KartuKredit struct {
	gorm.Model
	Tipe        string        `gorm:"type:varchar(50);not null" json:"tipe" form:"tipe"`
	Nama        string        `gorm:"type:varchar(255);not null" json:"nama" form:"nama"`
	CVV         int           `gorm:"not null" json:"cvv" form:"cvv"`
	Bulan       int           `gorm:"not null" json:"bulan" form:"bulan"`
	Tahun       int           `gorm:"not null" json:"tahun" form:"tahun"`
	Nomor       int           `gorm:"not null" json:"nomor" form:"nomor"`
	Reservation []Reservation `gorm:"foreignKey:KartuKreditID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
