package handlers

import (
	"database/sql"
	"log"
	"movies-api/entities"
	"movies-api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMovies(conn *sql.DB, c *gin.Context) {
	movies, err := repositories.GetAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, movies)
}

func GetMovieById(conn *sql.DB, c *gin.Context) {
	id := c.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "Param \"id\" must be a valid UUID."})
		return
	}
	movie, err := repositories.GetById(conn, parsedId)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, movie)
}

func PostMovie(conn *sql.DB, c *gin.Context) {
	var movieData entities.Movie

	if err := c.BindJSON(&movieData); err != nil {
		log.Printf("Error parsing body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "Invalid body."})
	}

	movie, err := repositories.Create(conn, movieData)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, movie)
}
