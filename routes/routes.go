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
	e.GET("/rooms", controllers.GetAllHomestayController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(echoMid.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	// Rooms JWT
	r.POST("/rooms", controllers.CreateHomestayController)
	r.GET("/rooms/:id", controllers.GetHomestayByIDController)
	r.PUT("/rooms/:id", controllers.UpdateHomestayController)
	r.DELETE("/rooms/:id", controllers.DeleteHomestayController)

	return e
}
