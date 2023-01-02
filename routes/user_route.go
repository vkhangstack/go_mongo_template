package routes

import (
	"go-mongo-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)     // create user
	e.POST("/auth/login", controllers.Login)    //
	e.POST("/auth/register", controllers.Login) //
}
