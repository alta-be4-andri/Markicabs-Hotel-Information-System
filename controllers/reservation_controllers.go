package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	"project2/response"

	"github.com/labstack/echo/v4"
)

func CreateReservationControllers(c echo.Context) error {
	input := models.Reservation{}
	c.Bind(&input)
	_, err := databases.CreateReservation(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
