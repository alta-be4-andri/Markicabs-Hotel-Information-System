package controllers

import (
	"io"
	"net/http"
	"os"
	"project2/response"

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
