package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/sample-api/api"
	"github.com/lottotto/sample-api/db"
	// "go.elastic.co/apm/module/apmechov4"
)

func main() {

	_, err := db.ConfigDB()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(apmechov4.Middleware())

	// Routes
	e.GET("/horse", api.GetHorsebyName)
	e.GET("/horse/:id", api.GetHorsebyID)
	e.GET("/race/:id", api.GetRaceByID)
	e.GET("/hc", api.Healthcheck)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
