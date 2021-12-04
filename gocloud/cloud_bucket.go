package gocloud

import (
	"io"
	"net/http"
	"net/url"
	"project2/lib/databases"
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

func HandleFileUploadToBucket(c echo.Context) error {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
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
	photo := models.Photo{
		RoomsID:    uint(id),
		Nama_Photo: sw.Attrs().Name,
		Url:        u.String(),
	}
	_, err = databases.InsertPhoto(&photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "file uploaded successfully",
		"pathname": photo.Url,
	})
}
