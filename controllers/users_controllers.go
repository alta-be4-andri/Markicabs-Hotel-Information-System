package controllers

import (
	"net/http"
	"project2/crypt"
	"project2/lib/databases"
	"project2/middlewares"
	"project2/models"
	"project2/response"

	"github.com/labstack/echo/v4"
)

var user models.Users

func CreateUserControllers(c echo.Context) error {
	c.Bind(&user)
	newPass, _ := crypt.Encrypt(user.Password)
	user.Password = newPass
	_, err := databases.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}

	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func GetUserControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := databases.GetUser(id)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData(user))
}

func UpdateUserControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	c.Bind(&user)
	_, err := databases.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

func DeleteUserControllers(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	_, err := databases.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}

//login users
func LoginUsersController(c echo.Context) error {
	user := models.UserLogin{}
	c.Bind(&user)
	users, err := databases.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.LoginFailedResponse())
	}
	return c.JSON(http.StatusOK, response.LoginSuccessResponse(users))
}
