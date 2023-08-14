package main

import (
	"log"
	"movies-api/config"
	"movies-api/database"
)

func main() {
	log.Printf("Config loaded:\n\tAPI: %v\n\tDB: %v\n", config.GetAPI(), config.GetDB())
	conn, err := database.OpenConnection()
	if err != nil {
		log.Fatalf("Error openning connection: %v\n", err)
	}

	log.Printf("Successfully opened database connection: %v", conn)
}
