package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lottotto/sample-api/api"
	// "go.elastic.co/apm/module/apmechov4"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(apmechov4.Middleware())

	// Routes
	e.GET("/horses/:id", api.GetHorsebyID)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
