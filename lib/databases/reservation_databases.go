package databases

import (
	"project2/config"
	"project2/models"
)

// Fungsi untuk membuat reservasi baru
func CreateReservation(reservation *models.Reservation) (interface{}, error) {
	tx := config.DB.Create(&reservation)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return reservation, nil
}

// Fungsi untuk mendapatkan reservasi by reservasi id
func GetReservation(id int) (interface{}, error) {
	var reservation models.Reservation
	tx := config.DB.Where("id = ?", id).Find(&reservation)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return reservation, nil
}

// Fungsi untuk mendapatkan reservasi owner
func GetReservationOwner(id int) (uint, error) {
	var reservation models.Reservation
	tx := config.DB.Where("id = ?", id).Find(&reservation)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return reservation.UsersID, nil
}

// Fungsi untuk menghapus reservasi by reservasi id
func CancelReservation(id int) (interface{}, error) {
	var reservation models.Reservation
	if err := config.DB.Where("id = ?", id).Delete(&reservation).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}
