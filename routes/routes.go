package routes

import (
	"project2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/users", controllers.CreateUserControllers)

	return e
}
