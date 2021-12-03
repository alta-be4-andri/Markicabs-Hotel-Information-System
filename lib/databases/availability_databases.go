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
	tx := config.DB.Table("reservations").Select("reservations.check_in, reservations.check_out").Where("reservations.rooms_id = ? AND reservations.check_out > ?", uint(id), time.Now()).Find(&dates)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return dates, nil
}

// Fungsi untuk mendapatkan jumlah malam pada reservasi
func HitungJumlahMalam(checkIn time.Time, checkOut time.Time, idReservation uint) {
	config.DB.Exec("UPDATE reservations SET jumlah_malam = (SELECT DATEDIFF(?, ?)) WHERE id = ?", checkOut, checkIn, idReservation)
}

// Fungsi untuk menambahkan harga pada reservasi
func GetHargaRoom(idRoom int) (int, error) {
	var harga int
	tx := config.DB.Raw("SELECT harga FROM rooms WHERE id = ?", idRoom).Scan(&harga)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return harga, nil
}
