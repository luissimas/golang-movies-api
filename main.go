package main

import (
	"log"
	"movies-api/config"
)

func main() {
	log.Printf("Config loaded:\n\tAPI: %v\n\tDB: %v\n", config.GetAPI(), config.GetDB())
}
