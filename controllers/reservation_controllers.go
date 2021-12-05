package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/response"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var format_date string = "2006-01-02"

// Controller untuk membuat reservasi baru
func CreateReservationControllers(c echo.Context) error {
	// Membuat reservasi baru
	body := models.ReservationBody{}
	c.Bind(&body)
	logged := middlewares.ExtractTokenUserId(c)

	input := models.Reservation{}

	input.Check_In, _ = time.Parse(format_date, body.Check_In)
	input.Check_Out, _ = time.Parse(format_date, body.Check_Out)
	if input.Check_In.Unix() > input.Check_Out.Unix() {
		return c.JSON(http.StatusBadRequest, response.DateInvalidResponse())
	}
	input.UsersID = uint(logged)
	input.RoomsID = body.RoomsID
	roomOwner, err := databases.GetRoomOwner(int(input.RoomsID))
	if err != nil || roomOwner == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if roomOwner == input.UsersID {
		return c.JSON(http.StatusBadRequest, response.AccessForbiddenResponse())
	}

	reservation, err := databases.CreateReservation(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	databases.AddJumlahMalam(input.Check_In, input.Check_Out, reservation.ID)
	totalHarga := databases.AddHargaToReservation(input.RoomsID, reservation.ID)

	var body2 = models.RequestBodyStruct{
		ReferenceID:    strconv.Itoa(int(reservation.ID)),
		Currency:       "IDR",
		Amount:         float64(totalHarga),
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_OVO",
		ChannelProperties: models.ChannelProperties{
			MobileNumber: body.Phone,
		},
		Metadata: models.Metadata{
			BranchArea: "PLUIT",
			BranchCity: "JAKARTA",
		},
	}

	reqBody, err := json.Marshal(body2)
	if err != nil {
		print(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.xendit.co/ewallets/charges", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.SetBasicAuth(os.Getenv("SECRET_KEY"), os.Getenv("PASS_XENDIT"))

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body4, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responsePayment models.ResponsePayment
	json.Unmarshal([]byte(body4), &responsePayment)
	// response
	return c.JSON(http.StatusOK, response.ReservationSuccessResponse(reservation.ID, responsePayment))

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
	getReservation, _ := databases.GetReservation(id)
	reservation := models.GetReservation{}
	reservation.RoomsID = getReservation.RoomsID
	reservation.Check_In = getReservation.Check_In.Format(format_date)
	reservation.Check_Out = getReservation.Check_Out.Format(format_date)
	reservation.Total_Harga = getReservation.Total_Harga
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
