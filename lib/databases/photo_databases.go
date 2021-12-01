package databases

import (
	"project2/config"
	"project2/models"
)

func InsertPhoto(photo *models.Photo) (interface{}, error) {
	if err := config.DB.Create(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

func GetAllPhoto() (interface{}, error) {
	var photo models.Photo
	var result []models.Get_Photo
	if err := config.DB.Model(&photo).Where("rooms_id = ?", photo.RoomsID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func DeletePhoto(id int) (interface{}, error) {
	var photo models.Photo
	if err := config.DB.Where("id = ?", id).Delete(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

func GetIDRoomsPhoto(id int) (uint, error) {
	var photo models.Photo
	err := config.DB.Find(&photo, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return uint(photo.RoomsID), nil
}
