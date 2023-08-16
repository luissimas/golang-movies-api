package main

import (
	"fmt"
	"log"
	"movies-api/config"
	"movies-api/database"
	"movies-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Config loaded:\n\tAPI: %v\n\tDB: %v\n", config.GetAPI(), config.GetDB())
	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatalf("Error creating database connection: %v\n", err)
	}
	log.Print("Connected to the database.")

	router := gin.Default()

	router.GET("/movie", handlers.AdaptHandler(handlers.GetMovies, db))
	router.GET("/movie/:id", handlers.AdaptHandler(handlers.GetMovieById, db))
	router.POST("/movie", handlers.AdaptHandler(handlers.PostMovie, db))
	router.DELETE("/movie/:id", handlers.AdaptHandler(handlers.DeleteMovie, db))

	url := fmt.Sprintf("localhost:%s", config.GetAPI().Port)

	router.Run(url)
}
