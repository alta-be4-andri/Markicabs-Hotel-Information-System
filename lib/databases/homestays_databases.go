package databases

import (
	"project2/config"
	"project2/models"
)

var homestay *models.HomeStay

// Fungsi untuk membuat tempat penyewaan homestay baru
func CreateHomestay(homestay *models.HomeStay) (*models.HomeStay, error) {
	if err := config.DB.Create(&homestay).Error; err != nil {
		return nil, err
	}
	return homestay, nil
}

func GetAllHomestays() (interface{}, error) {
	var results []models.Get_HomeStay
	if err := config.DB.Table("home_stays").Select("home_stays.id, home_stays.nama, home_stays.users_id, home_stays.longitude, home_stays.latitude, home_stays.alamat, home_stays.rating, home_stay_photos.url").Joins(
		"join home_stay_photos on home_stays.id = home_stay_photos.home_stay_id").Where("home_stays.deleted_at IS NULL").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func GetHomestaysByID(id int) (interface{}, error) {
	var result models.Get_HomeStay
	err := config.DB.Table("home_stays").Select("home_stays.id, home_stays.nama, home_stays.users_id, home_stays.longitude, home_stays.latitude, home_stays.alamat, home_stays.rating, home_stay_photos.url").Joins(
		"join home_stay_photos on home_stays.id = home_stay_photos.home_stay_id").Where("home_stays.deleted_at IS NULL AND home_stays.id =?", id).Find(&result)
	rows_affected := err.RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return result, nil
}

func GetHomestaysByKotaId(id int) (interface{}, error) {
	var result []models.Get_HomeStay
	err := config.DB.Table("home_stays").Where("kota_id = ? AND home_stays.deleted_at IS NULL", id).Find(&result)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected < 1 {
		return nil, nil
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

func GetIDUserHomestay(id int) (uint, error) {
	err := config.DB.Find(&homestay, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return homestay.UsersID, nil
}

func GetKota(id int) (string, error) {
	var town models.Kota
	if err := config.DB.Where("id = ?", id).Find(&town); err.Error != nil {
		return "", err.Error
	}
	return town.Nama_Kota, nil
}

func GetRating(id int) (int, error) {
	var rating models.HomeStay
	if err := config.DB.Where("id = ?", id).Find(&rating).Error; err != nil {
		return 0, err
	}
	return int(rating.Rating), nil
}

func AverageRatings(rating int) (float64, error) {
	var review models.HomeStay
	if err := config.DB.Raw("SELECT AVG(rating) FROM reviews").Scan(&review).Error; err != nil {
		return 0, err
	}
	return review.Rating, nil
}

func GetUrlHomestayPhoto(id int) (string, error) {
	var url string
	if err := config.DB.Table("home_stay_photos").Select("url").Where("id = ?", id).Find(&url).Error; err != nil {
		return "", err
	}
	return url, nil
}
