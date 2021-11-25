package routes

import (
	"project2/constant"
	"project2/controllers"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.POST("/users", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUsersController)

	r := e.Group("/jwt")
	r.Use(echoMid.JWT([]byte(constant.SECRET_JWT)))
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	return e
}
