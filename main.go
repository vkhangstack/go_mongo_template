package main

import (
	"go-mongo-api/configs"
	"go-mongo-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// initialize the database
	configs.ConnectDB()
	// routes
	routes.UserRoute(e)

	e.GET("/dev", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello backend"})
	})

	e.Logger.Fatal(e.Start(":4000"))
}
