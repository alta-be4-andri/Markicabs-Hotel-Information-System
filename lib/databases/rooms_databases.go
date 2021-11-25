package databases

import (
	"project2/config"
	"project2/models"
)

var result models.Get_Rooms

func CreateRoom(room *models.Rooms) (interface{}, error) {
	if err := config.DB.Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func GetAllRooms(room *models.Rooms) (interface{}, error) {
	if err := config.DB.Model(room).Find(result).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func GetRoomsByID(id int, room *models.Rooms) (interface{}, error) {
	err := config.DB.Model(room).Find(result, id)
	rows_affected := err.RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return result, nil
}

func UpdateRooms(id int, room *models.Rooms) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func DeleteRooms(id int, room *models.Rooms) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
