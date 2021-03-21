package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "go.elastic.co/apm/module/apmechov4"
)

// User ...
type User struct {
	Name string `json:"name"`
}

type Horst struct {
	Name    string `json:"name"`
	Trainer string `json:"trainer"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(apmechov4.Middleware())

	// Routes
	e.GET("/users/:name", hello)
	e.GET("/horses/:id", controller.getHorsebyID)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	name := c.Param("name")
	user := &User{Name: name}
	return c.JSON(http.StatusOK, user)
}
