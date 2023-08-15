package main

import (
	"database/sql"
	"fmt"
	"log"
	"movies-api/config"
	"movies-api/database"
	"movies-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Config loaded:\n\tAPI: %v\n\tDB: %v\n", config.GetAPI(), config.GetDB())
	conn, err := database.OpenConnection()
	if err != nil {
		log.Fatalf("Error creating database connection: %v\n", err)
	}
	log.Print("Connected to the database.")

	router := gin.Default()

	adaptHandler := func(handler func(*sql.DB, *gin.Context), conn *sql.DB) func(*gin.Context) {
		return func(ctx *gin.Context) {
			handler(conn, ctx)
		}
	}

	router.GET("/movie", adaptHandler(handlers.GetMovies, conn))
	router.GET("/movie/:id", adaptHandler(handlers.GetMovieById, conn))
	router.POST("/movie", adaptHandler(handlers.PostMovie, conn))

	url := fmt.Sprintf("localhost:%s", config.GetAPI().Port)

	router.Run(url)
}
