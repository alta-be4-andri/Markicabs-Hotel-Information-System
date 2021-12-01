package controllers

import (
	"io"
	"net/http"
	"os"
	"project2/lib/databases"
	"project2/models"
	"project2/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PhotoControllers(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}

		des, err := os.Create(file.Filename)
		if err != nil {
			return err
		}

		if _, err := io.Copy(des, src); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func InsertPhotoController(c echo.Context) error {
	var photo models.Photo
	c.Bind(photo)
	_, err := databases.InsertPhoto(&photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetAllPhotoController(c echo.Context) error {
	photo, err := databases.GetAllPhoto()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(photo))
}

func UpdatePhotoController(c echo.Context) error {
	var photo models.Photo
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	c.Bind(&homestay)
	_, err = databases.UpdatePhoto(id, &photo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeletePhotoController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	_, err = databases.DeletePhoto(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
