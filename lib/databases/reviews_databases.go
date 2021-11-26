package databases

import (
	"project2/config"
	"project2/models"
)

func AddReviews(review *models.Review) (interface{}, error) {
	if err := config.DB.Create(&review).Error; err != nil {
		return err, nil
	}
	return review, nil
}

func GetReviews(id int) (interface{}, error) {
	var results []models.Get_Review
	tx := config.DB.Table("reviews").Where("home_stay_id = ?", id).Scan(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &results, nil
}

func AddRatingToHomestay(id int) {
	config.DB.Exec("UPDATE home_stays SET rating = (SELECT AVG(rating) FROM reviews WHERE home_stay_id = ?) WHERE id = ?", id, id)
}
