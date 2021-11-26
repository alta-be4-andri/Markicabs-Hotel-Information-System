package controllers

import (
	"net/http"
	"project2/lib/databases"
	"project2/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type BodyCheckIn struct {
	Check_In  time.Time `json:"check_in" form:"check_in"`
	Check_Out time.Time `json:"check_out" form:"check_out"`
}

// Fungsi untuk melakukan pengecekan availability suatu room
func RoomReservationCheck(c echo.Context) error {
	input := BodyCheckIn{}
	c.Bind(&input)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	// Mendapatkan seluruh tanggal reservation room tertentu
	dateList, err := databases.RoomReservationList(id)
	if err != nil || dateList == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	// Pengecekan ketersediaan room untuk tanggal check_in dan check_out yang diinginkan
	for _, date := range dateList {
		input_checkin := input.Check_In.Unix()
		input_checkout := input.Check_Out.Unix()
		date_checkin := date.Check_In.Unix()
		date_checkout := date.Check_Out.Unix()
		if (input_checkin >= date_checkin && input_checkin <= date_checkout) || (input_checkout >= date_checkin && input_checkout <= date_checkout) {
			return c.JSON(http.StatusBadRequest, response.CheckFailedResponse())
		}
	}
	return c.JSON(http.StatusBadRequest, response.CheckSuccessResponse())
}