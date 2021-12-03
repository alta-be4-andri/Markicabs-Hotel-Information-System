package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/plugins"
	"project2/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

var homestay models.HomeStay

func CreateHomestayController(c echo.Context) error {
	c.Bind(&homestay)
	logged := middlewares.ExtractTokenUserId(c)
	homestay.UsersID = uint(logged)
	get_kota, _ := databases.GetKota(homestay.KotaID)
	lat, long, _ := plugins.Geocode(get_kota)
	homestay.Latitude = lat
	homestay.Longitude = long
	_, err := databases.CreateHomestay(&homestay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetAllHomestayController(c echo.Context) error {
	homestay, err := databases.GetAllHomestays()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.HomestayNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(homestay))
}

func GetHomestayByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	homestay, err := databases.GetHomestaysByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.HomestayNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(homestay))
}

func GetHomestayByKotaIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	homestay, err := databases.GetHomestaysByKotaId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if homestay == nil {
		return c.JSON(http.StatusBadRequest, response.HomestayNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(homestay))
}

func UpdateHomestayController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != id_user_homestay {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	c.Bind(&homestay)
	databases.UpdateHomestays(id, &homestay)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeleteHomestayController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	id_user_homestay, _ := databases.GetIDUserHomestay(id)
	logged := middlewares.ExtractTokenUserId(c)
	if uint(logged) != id_user_homestay {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}
	databases.DeleteHomestays(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
