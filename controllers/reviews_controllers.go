package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	"project2/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

var review models.Review

func AddReviewsController(c echo.Context) error {
	c.Bind(&review)

	_, err := databases.AddReviews(&review)
	databases.AddRatingToHomestay(int(review.HomeStayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetReviewsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	review, err := databases.GetReviews(id)
	if err != nil || review == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(&review))
}
