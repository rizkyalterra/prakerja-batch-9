package routes

import (
	"km5go/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))

	eLog := eAuth.Group("")
	eLog.Use(middleware.Logger())
	eLog.GET("/news", controllers.GetNewsController)
	
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.POST("/news", controllers.AddNewsController)
	eAuth.GET("/news/:id", controllers.GetDetailNewsController)
}