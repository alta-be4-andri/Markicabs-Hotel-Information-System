package databases

import (
	"project2/config"
	"project2/models"
)

// Fungsi untuk membuat tempat room baru
func CreateRoom(room *models.BodyRoom) (*models.Rooms, error) {
	input := models.Rooms{
		Nama_Room:      room.Nama_Room,
		HomeStayID:     room.HomeStayID,
		Total_Penghuni: room.Total_Penghuni,
		Harga:          room.Harga,
		Deskripsi:      room.Deskripsi,
	}
	if err := config.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

// Fungsi untuk memasukkan fasilitas di room
func CreateRoomFasilitas(idRoom uint, idFasilitas int) (interface{}, error) {
	input := models.FasilitasRoom{
		RoomsID:     idRoom,
		FasilitasID: uint(idFasilitas),
	}
	if err := config.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return input, nil
}

//Fungsi untuk mendapatkan semua room
func GetAllRooms() (interface{}, error) {
	var results []models.AllRoomsResponse
	tx := config.DB.Table("rooms").Select(
		"rooms.id, rooms.home_stay_id,rooms.nama_room,rooms.total_penghuni, rooms.harga, rooms.deskripsi, room_photos.url").Joins(`
		"join room_photos on rooms.id = room_photos.rooms_id"`).Where("rooms.deleted_at IS NULL").Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return results, nil
}

//Fungsi untuk mendapatkan room by homestay id
func GetRoomByHomestayID(id int) (interface{}, error) {
	var results []models.AllRoomsResponse
	tx := config.DB.Table("rooms").Select(
		"rooms.id, rooms.home_stay_id,rooms.nama_room,rooms.total_penghuni, rooms.harga, rooms.deskripsi, room_photos.url").Joins(`
		"join room_photos on rooms.id = room_photos.rooms_id"`).Where("rooms.home_stay_id = ? AND rooms.deleted_at IS NULL", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan room by id
func GetRoomByID(id int) (*models.FasilitasRoomResponse, error) {
	var results models.FasilitasRoomResponse
	tx := config.DB.Table("rooms").Select(
		"rooms.id,rooms.nama_room, rooms.home_stay_id,rooms.total_penghuni, rooms.harga, rooms.deskripsi, room_photos.url").Joins(`
		"join room_photos on rooms.id = room_photos.rooms_id"`).Where("rooms.id = ? AND rooms.deleted_at IS NULL", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &results, nil
}

// Fungsi untuk mendapatkan fasilitas room tertentu
func GetFasilitasRoom(id int) ([]string, error) {
	var results []string
	tx := config.DB.Table("fasilitas_rooms").Select(
		"fasilitas.nama_fasilitas").Joins("join fasilitas on fasilitas.id = fasilitas_rooms.fasilitas_id").Where("fasilitas_rooms.rooms_id = ?", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

//Fungsi untuk mendapatkan pemilik room
func GetRoomOwner(id int) (uint, error) {
	var result uint
	tx := config.DB.Table("rooms").Select(
		"home_stays.users_id").Joins("join home_stays on rooms.home_stay_id = home_stays.id").Where("rooms.id = ?", id).Find(&result)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return result, nil
}

// Fungsi untuk melakukan update room
func UpdateRoom(id int, room *models.Rooms) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func DeleteRoom(id int) (interface{}, error) {
	var room models.Rooms
	if err := config.DB.Where("id = ?", id).Delete(&room).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}
