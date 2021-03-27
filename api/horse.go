package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lottotto/db"
	"github.com/lottotto/model"
)

func GetHorsebyID(c echo.Context) error {
	var horse *model.Horse
	rows, err := db.Query("SELECT * from horse where id=?", c.Param("id"))
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&horse.Id, &horse.Name, &horse.Age, &horse.Sex, &horse.Trainer, &horse.Owner)
		if err != nil {
			log.Fatal(err)
		}
	}

	return c.String(http.StatusOK, horse)
}
