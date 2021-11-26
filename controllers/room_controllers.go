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

func GetAllRoomsController(c echo.Context) error {
	room, err := databases.GetAllRooms()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(room))
}

func GetRoomByHomestayIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	room, err := databases.GetRoomByHomestayID(id)
	if err != nil || room == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(room))
}

func GetRoomByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	room, err := databases.GetRoomByID(id)
	if err != nil || room == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	room.Fasilitas, _ = databases.GetFasilitasRoom(id)
	return c.JSON(http.StatusOK, response.SuccessResponseData(room))
}

func UpdateRoomController(c echo.Context) error {
	var room models.Rooms
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_room, _ := databases.GetRoomOwner(id)
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != id_user_room {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	c.Bind(&room)
	databases.UpdateRoom(id, &room)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeleteRoomController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_room, _ := databases.GetRoomOwner(id)
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != id_user_room {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteRoom(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
