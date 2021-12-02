package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type BodyDate struct {
	Check_In  string `json:"check_in" form:"check_in"`
	Check_Out string `json:"check_out" form:"check_out"`
}

type InputDate struct {
	Check_In  time.Time `json:"check_in" form:"check_in"`
	Check_Out time.Time `json:"check_out" form:"check_out"`
}

// Fungsi untuk melakukan pengecekan availability suatu room
func RoomReservationCheck(c echo.Context) error {
	body := BodyDate{}
	c.Bind(&body)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	// Mendapatkan seluruh tanggal reservation room tertentu
	dateList, err := databases.RoomReservationList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	input := InputDate{}
	input.Check_In, _ = time.Parse(format_date, body.Check_In)
	input.Check_Out, _ = time.Parse(format_date, body.Check_Out)

	if input.Check_In.Unix() < time.Now().Unix() || input.Check_Out.Unix() < time.Now().Unix() {
		return c.JSON(http.StatusBadRequest, response.CheckFailedResponse())
	}

	// Pengecekan ketersediaan room untuk tanggal check_in dan check_out yang diinginkan
	for _, date := range dateList {
		if (input.Check_In.Unix() >= date.Check_In.Unix() && input.Check_In.Unix() <= date.Check_Out.Unix()) || (input.Check_Out.Unix() >= date.Check_In.Unix() && input.Check_Out.Unix() <= date.Check_Out.Unix()) {
			return c.JSON(http.StatusBadRequest, response.CheckFailedResponse())
		}
	}
	return c.JSON(http.StatusBadRequest, response.CheckSuccessResponse())
}
