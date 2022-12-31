package routes

import (
	"go-mongo-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser) // create user
}
