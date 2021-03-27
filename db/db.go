package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

type DB struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	Connection *sql.DB
}

var d *DB

func configDB() *DB {
	d = &DB{Host: getEnv("POSTGRES_HOST", "127.0.0.1"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		Username: getEnv("POSTGERS_USER", "postgres"),
		Password: getEnv("POSTGRES_PASS", "postgres")}
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.Username, d.Password, d.DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(fmt.Sprintf("Connection Faild: %s", dbinfo))
		panic(err)
	}
	d.Connection = db
	return d
}
