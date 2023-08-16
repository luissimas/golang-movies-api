package handlers

import (
	"errors"
	"fmt"
	"log"
	. "movies-api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AdaptHandler(handler func(*gorm.DB, *gin.Context), db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		handler(db, c)
	}
}

func GetMovies(db *gorm.DB, c *gin.Context) {
	var movies []Movie
	err := db.Find(&movies).Error
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, movies)
}

func GetMovieById(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "Param \"id\" must be a valid UUID."})
		return
	}
	var movie = &Movie{ID: parsedId}
	err = db.First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, map[string]string{"message": "Movie not found."})
			return
		}
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, movie)
}

func PostMovie(db *gorm.DB, c *gin.Context) {
	var movie Movie

	if err := c.BindJSON(&movie); err != nil {
		log.Printf("Error parsing body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "Invalid body."})
	}

	if err := db.Create(&movie).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("movie: %v+", movie)
	c.IndentedJSON(http.StatusCreated, movie)
}

func DeleteMovie(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "Param \"id\" must be a valid UUID."})
		return
	}

	var movie = &Movie{ID: parsedId}
	err = db.First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, map[string]string{"message": "Movie not found."})
			return
		}
		log.Fatal(err)
	}

	if err := db.Delete(&movie).Error; err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
