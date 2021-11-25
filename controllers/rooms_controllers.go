package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/models"
	"project2/response"

	"github.com/labstack/echo/v4"
)

var rooms models.Rooms

func CreateRoomsController(c echo.Context) error {
	c.Bind(rooms)

	_, err := databases.CreateRoom(&rooms)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
