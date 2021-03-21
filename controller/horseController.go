package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getHorsebyID(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
