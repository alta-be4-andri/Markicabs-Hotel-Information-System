package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UsersID       uint      `json:"users_id" form:"users_id"`
	RoomsID       uint      `json:"rooms_id" form:"rooms_id"`
	Check_In      time.Time `gorm:"type:datetime;not null" json:"check_in" form:"check_in"`
	Jumlah_Malam  int       `json:"jumlah_malam" form:"jumlah_malam"`
	Check_Out     time.Time `gorm:"type:datetime" json:"check_out" form:"check_out"`
	KartuKreditID uint      `gorm:"default:0" json:"kartu_kredit_id" form:"kartu_kredit_id"`
	Total_Harga   int       `json:"total_harga" form:"total_harga"`
	Review        Review    `gorm:"foreignKey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
