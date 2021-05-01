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

var D *DB

func ConfigDB() (*DB, error) {
	D = &DB{Host: getEnv("POSTGRES_HOST", "127.0.0.1"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		Username: getEnv("POSTGERS_USER", "postgres"),
		Password: getEnv("POSTGRES_PASS", "postgres"),
		DBName:   getEnv("POSTGRES_NAME", "sample"),
	}
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", D.Host, D.Port, D.Username, D.Password, D.DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(fmt.Sprintf("Connection Faild: %s", dbinfo))
		return nil, err
	}
	D.Connection = db
	return D, nil
}
