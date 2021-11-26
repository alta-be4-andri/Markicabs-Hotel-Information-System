package databases

import (
	"project2/config"
	"time"
)

type ReservationDate struct {
	Check_In  time.Time
	Check_Out time.Time
}

// Fungsi untuk mendapatkan tanggal check_in dan check_out suatu reservasi
func RoomReservationList(id int) ([]ReservationDate, error) {
	var dates []ReservationDate
	tx := config.DB.Table("reservations").Select("reservations.check_in, reservations.check_out").Where("reservations.rooms_id = ?", uint(id)).Find(&dates)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return dates, nil
}
