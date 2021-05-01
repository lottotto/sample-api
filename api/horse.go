package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lottotto/sample-api/db"
	"github.com/lottotto/sample-api/model"
)

func GetHorsebyName(c echo.Context) error {
	var horse model.Horse
	conn := db.D.Connection
	name := c.Param("name")
	// query := fmt.Sprintf("SELECT * from horse where name=%s", name)
	// fmt.Println(query)
	err := conn.QueryRow("SELECT * from horse where name=$1", name).Scan(&horse.Id, &horse.Name, &horse.Trainer, &horse.Owner, &horse.Breeder, &horse.Sire, &horse.Broodmare)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, horse)
}

func GetHorsebyID(c echo.Context) error {
	var horse model.Horse
	conn := db.D.Connection
	rows, err := conn.Query("SELECT * from horse")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&horse.Id, &horse.Name, &horse.Trainer, &horse.Owner, &horse.Breeder, &horse.Sire, &horse.Broodmare)
		if err != nil {
			log.Fatal(err)
		}
		if horse.Id == c.Param("id") {
			return c.JSON(http.StatusOK, horse)
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}
