package routes

import (
	"km5go/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersController)
	e.GET("/news", controllers.GetNewsController)
	e.POST("/news", controllers.AddNewsController)
	e.GET("/news/:id", controllers.GetDetailNewsController)
}