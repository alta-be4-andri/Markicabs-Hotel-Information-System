package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	RoomsID    int
	Nama_Photo string
	Url        string
}

type Get_Photo struct {
	RoomsID    int
	Nama_Photo string
}
