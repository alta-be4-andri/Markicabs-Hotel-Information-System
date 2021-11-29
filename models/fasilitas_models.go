package models

type Fasilitas struct {
	ID             int             `gorm:"primaryKey; autoIncrement" json:"id" form:"id"`
	Nama_Fasilitas string          `gorm:"type:varchar(255);not null" json:"nama_fasilitas" form:"nama_fasilitas"`
	FasilitasRoom  []FasilitasRoom `gorm:"foreignKey:FasilitasID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
