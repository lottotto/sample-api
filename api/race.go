package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lottotto/sample-api/db"
	"github.com/lottotto/sample-api/model"
)

func GetRaceByID(c echo.Context) error {
	var race model.Race
	conn := db.D.Connection
	rows, err := conn.Query("SELECT * from race")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&race.Id, &race.Name, &race.Distance, &race.Field, &race.Rotation, &race.Condition, &race.Course)
		if err != nil {
			log.Fatal(err)
		}
		if race.Id == c.Param("id") {
			return c.JSON(http.StatusOK, race)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}
