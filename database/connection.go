package database

import (
	"database/sql"
	"fmt"
	"log"
	"movies-api/config"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDB()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}

	err = conn.Ping()

	return conn, err
}
