package databases

import (
	"project2/config"
	"project2/models"
)

// var reviewResponse models.Get_Review

func AddReviews(review *models.Review) (interface{}, error) {
	if err := config.DB.Create(&review).Error; err != nil {
		return err, nil
	}
	return review, nil
}

// func GetReviews(id int) (float64, error) {
// 	if err := config.DB.Exec("UPDATE home_stays SET rating = (SELECT AVG(rating) FROM reviews WHERE home_stays_id = ?) WHERE ID = ?", homestay.ID).Error; err != nil {
// 		return 0, nil
// 	}
// 	return homestay.Rating, nil
// }

func AddRatingToHomestay(id int) {
	config.DB.Exec("UPDATE home_stays SET rating = (SELECT AVG(rating) FROM reviews WHERE home_stays_id = ?) WHERE ID = ?", id)
}
