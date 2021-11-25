package databases

import (
	"project2/config"
	"project2/models"
)

var result models.Get_HomeStay
var homestay *models.HomeStay

func CreateHomestay(homestay *models.HomeStay) (interface{}, error) {
	if err := config.DB.Create(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}

func GetAllHomestays() (interface{}, error) {
	var results []models.Get_HomeStay
	var homestays []models.HomeStay
	if err := config.DB.Model(homestays).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func GetHomestaysByID(id int) (interface{}, error) {
	err := config.DB.Model(&homestay).Find(&result, id)
	rows_affected := err.RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return result, nil
}

func UpdateHomestays(id int, homestay *models.HomeStay) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}

func DeleteHomestays(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}

// function bantuan untuk mendapatkan id user pada tabel produk
func GetIDUserHomestay(id int) (uint, error) {
	err := config.DB.Find(&homestay, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return homestay.UsersID, nil
}
