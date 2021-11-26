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
	e.GET("/homestays", controllers.GetAllHomestayController)
	e.GET("/rooms", controllers.GetAllRoomsController)
	e.GET("/rooms/homestays/:id", controllers.GetRoomByHomestayIdController)
	e.GET("/rooms/:id", controllers.GetRoomByIdController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(echoMid.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	// Homestay JWT
	r.POST("/homestays", controllers.CreateHomestayController)
	r.GET("/homestays/:id", controllers.GetHomestayByIDController)
	r.PUT("/homestays/:id", controllers.UpdateHomestayController)
	r.DELETE("/homestays/:id", controllers.DeleteHomestayController)

	// Room JWT
	r.PUT("/rooms/:id", controllers.UpdateRoomController)
	r.DELETE("/rooms/:id", controllers.DeleteRoomController)

	return e
}
