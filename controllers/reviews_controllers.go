package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	"project2/response"

	"github.com/labstack/echo/v4"
)

var review models.Review

func AddReviewsController(c echo.Context) error {
	c.Bind(&review)

	_, err := databases.AddReviews(&review)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	// databases.AddRatingToHomestay(int(review.HomeStayID))
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

// func GetReviewsController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
// 	}
// 	if err != nil || review == 0 {
// 		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
// 	}
// 	return c.JSON(http.StatusOK, response.SuccessResponseData(&review))
// }
