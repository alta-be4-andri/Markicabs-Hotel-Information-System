package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReservationID uint   `json:"reservation_id" form:"reservation_id"`
	Rating        int    `gorm:"type:enum('1','2','3','4','5');not null" json:"rating" form:"rating"`
	Comment       string `gorm:"type:longtext" json:"comment" form:"comment"`
	HomeStayID    uint   `json:"homestay_id" form:"homestay_id"`
}

type Get_Review struct {
	ID      uint
	Rating  int
	Comment string
}
