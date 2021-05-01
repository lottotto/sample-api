package main

import (
	"github.com/lottotto/sample-api/db"
	"github.com/lottotto/sample-api/route"
	// "go.elastic.co/apm/module/apmechov4"
)

func main() {

	_, err := db.ConfigDB()
	if err != nil {
		panic(err)
	}
	// create echo router
	e := route.Init()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
