package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/response"

	"github.com/labstack/echo/v4"
)

func InputDataFasilitasCon(c echo.Context) error {
	input := databases.InputData()
	if input != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
