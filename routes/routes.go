package routes

import (
	"project2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/users", controllers.CreateUserControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.POST("/login", controllers.LoginUsersController)

	return e
}
