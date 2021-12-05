package controllers

import (
	"io"
	"net/http"
	"net/url"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/response"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func CreateRoomController(c echo.Context) error {
	body := models.BodyRoom{}
	c.Bind(&body)
	idHomestay, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	body.HomeStayID = uint(idHomestay)
	room, err := databases.CreateRoom(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	for _, fasilitas := range body.Fasilitas {
		databases.CreateRoomFasilitas(room.ID, fasilitas)
	}
	bucket := "project2-airbnb" //your bucket name

	ctx := appengine.NewContext(c.Request())

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	sw := storageClient.Bucket(bucket).Object(file.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"Error":   true,
		})
	}
	photo := models.RoomPhoto{
		RoomsID:    uint(room.ID),
		Nama_Photo: sw.Attrs().Name,
		Url:        u.String(),
	}
	_, err = databases.InsertRoomPhoto(&photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Room created and file uploaded successfully",
		"pathname": photo.Url,
	})
}

func GetAllRoomsController(c echo.Context) error {
	room, err := databases.GetAllRooms()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if room == nil {
		return c.JSON(http.StatusBadRequest, response.RoomNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(room))
}

func GetRoomByHomestayIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	room, err := databases.GetRoomByHomestayID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if room == nil {
		return c.JSON(http.StatusBadRequest, response.RoomNotFoundResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(room))
}

func GetRoomByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	room, err := databases.GetRoomByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	if room == nil {
		return c.JSON(http.StatusBadRequest, response.RoomNotFoundResponse())
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
