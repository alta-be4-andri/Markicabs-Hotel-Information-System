package routes

import (
	"net/http"
	"project2/constant"
	"project2/controllers"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(echoMid.CORSWithConfig(echoMid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/input", controllers.InputDataFasilitasCon)
	e.POST("/signup", controllers.CreateUserControllers)
	e.POST("/signin", controllers.LoginUsersController)
	e.GET("/homestays", controllers.GetAllHomestayController)
	e.GET("/homestays/:id", controllers.GetHomestayByIDController)
	e.GET("/homestays/kota/:id", controllers.GetHomestayByKotaIdController)
	e.GET("/rooms", controllers.GetAllRoomsController)
	e.GET("/rooms/homestays/:id", controllers.GetRoomByHomestayIdController)
	e.GET("/rooms/:id", controllers.GetRoomByIdController)
	e.GET("/rooms/check/:id", controllers.RoomReservationCheck)
	e.GET("/reviews/:id", controllers.GetReviewsController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(echoMid.JWT([]byte(constant.SECRET_JWT)))

	// Users JWT
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users", controllers.UpdateUserControllers)
	r.DELETE("/users", controllers.DeleteUserControllers)

	// Homestay JWT
	r.POST("/homestays", controllers.CreateHomestayController)
	r.PUT("/homestays/:id", controllers.UpdateHomestayController)
	r.DELETE("/homestays/:id", controllers.DeleteHomestayController)

	// Room JWT
	r.POST("/rooms/:id", controllers.CreateRoomController)
	r.PUT("/rooms/:id", controllers.UpdateRoomController)
	r.DELETE("/rooms/:id", controllers.DeleteRoomController)
	r.POST("/rooms", controllers.InsertPhotoController)
	r.GET("/rooms", controllers.GetAllPhotoController)
	r.GET("/rooms/:id", controllers.DeletePhotoController)

	// Review JWT
	r.POST("/reviews", controllers.AddReviewsController)

	// Reservation JWT
	r.POST("/reservations", controllers.CreateReservationControllers)
	r.GET("/reservations/:id", controllers.GetReservationControllers)
	r.DELETE("/reservations/:id", controllers.CancelReservationController)

	return e
}
