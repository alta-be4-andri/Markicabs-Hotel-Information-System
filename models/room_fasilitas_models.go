package models

import "gorm.io/gorm"

type FasilitasRoom struct {
	gorm.Model
	RoomsID     uint `json:"rooms_id" form:"rooms_id"`
	FasilitasID uint `json:"fasilitas_id" form:"fasilitas_id"`
}
