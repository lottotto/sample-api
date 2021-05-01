package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lottotto/sample-api/db"
	"github.com/lottotto/sample-api/model"
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
	e.GET("/horse", GetHorsebyName)
	e.GET("/horse/:id", GetHorsebyID)
	e.GET("/race/:id", GetRaceByID)
	e.GET("/hc", HealthCheck)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO WORLD")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

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
