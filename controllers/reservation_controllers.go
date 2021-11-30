package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateReservationControllers(c echo.Context) error {
	input := models.Reservation{}
	c.Bind(&input)
	logged := middlewares.ExtractTokenUserId(c)
	input.UsersID = uint(logged)
	reservation, err := databases.CreateReservation(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	databases.AddJumlahMalam(input.Check_In, input.Check_Out, reservation.ID)
	databases.AddHargaToReservation(input.RoomsID, reservation.ID)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetReservationControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	userId, _ := databases.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != userId {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	reservation, _ := databases.GetReservation(id)
	return c.JSON(http.StatusOK, response.SuccessResponseData(reservation))
}

func CancelReservationController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	userId, _ := databases.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != userId {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.CancelReservation(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
