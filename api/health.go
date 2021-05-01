package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lottotto/sample-api/db"
	"github.com/lottotto/sample-api/model"
)

func HealthCheck(c echo.Context) error {
	h := &model.Health{}
	err := db.D.Connection.Ping()
	if err != nil {
		h.Status = http.StatusInternalServerError
		h.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, h)
	}
	h.Status = http.StatusOK
	return c.JSON(http.StatusOK, h)
}

func Hc() echo.HandlerFunc {
	return func(c echo.Context) error {
		h := &model.Health{}
		err := db.D.Connection.Ping()
		if err != nil {
			h.Status = http.StatusInternalServerError
			h.Message = err.Error()
			return c.JSON(http.StatusInternalServerError, h)
		}
		h.Status = http.StatusOK
		return c.JSON(http.StatusOK, h)
	}
}
