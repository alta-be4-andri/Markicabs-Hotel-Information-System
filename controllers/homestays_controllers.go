package controllers

import (
	"io"
	"net/http"
	"net/url"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/plugins"
	"project2/response"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func CreateHomestayController(c echo.Context) error {
	var storageClient *storage.Client
	var homestay models.HomeStay
	c.Bind(&homestay)
	logged := middlewares.ExtractTokenUserId(c)
	homestay.UsersID = uint(logged)
	get_kota, _ := databases.GetKota(homestay.KotaID)
	lat, long, _ := plugins.Geocode(get_kota)
	homestay.Latitude = lat
	homestay.Longitude = long
	createdHomestay, err := databases.CreateHomestay(&homestay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	bucket := "project2-airbnb-homestays" //your bucket name

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
	photo := models.HomeStayPhoto{
		HomeStayID: uint(createdHomestay.ID),
		Nama_Photo: sw.Attrs().Name,
		Url:        u.String(),
	}
	_, err = databases.InsertHomestayPhoto(&photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "homestay created and file uploaded successfully",
		"pathname": photo.Url,
	})
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
	var homestay models.HomeStay
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
