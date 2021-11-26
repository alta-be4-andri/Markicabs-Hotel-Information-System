package databases

import (
	"project2/config"
	"project2/models"
)

func CreateReservation(reservation *models.Reservation) (interface{}, error) {
	tx := config.DB.Create(&reservation)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return reservation, nil
}

func GetReservation(id int) (interface{}, error) {
	var reservation []models.Reservation
	tx := config.DB.Table("reservation").Where("users_id = ?", id).Find(&reservation)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return reservation, nil
}

func CancelReservation(id int) (interface{}, error) {
	var reservation models.Reservation
	if err := config.DB.Where("id = ?", id).Delete(&reservation).Error; err != nil {
		return nil, err
	}
	return "deleted", nil

}
