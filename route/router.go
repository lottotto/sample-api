package route

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/sample-api/api"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(apmechov4.Middleware())

	// Routes
	e.GET("/horse", api.GetHorsebyName)
	e.GET("/horse/:id", api.GetHorsebyID)
	e.GET("/race/:id", api.GetRaceByID)
	e.GET("/hc", api.HealthCheck)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD")
	})
	return e
}
