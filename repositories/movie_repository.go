package repositories

import (
	"database/sql"
	"fmt"
	"movies-api/entities"

	"github.com/google/uuid"
)

func Create(conn *sql.DB, movie entities.Movie) (entities.Movie, error) {
	id := uuid.New()
	_, err := conn.Exec("INSERT INTO movies (id, name, release_date, director, description, duration, budget) VALUES ($1, $2, $3, $4, $5, $6, $7)", id, movie.Name, movie.ReleaseDate, movie.Director, movie.Description, movie.Duration, movie.Budget)
	movie.ID = id
	if err != nil {
		return movie, fmt.Errorf("Create: %v", err)
	}
	return movie, nil
}

func GetAll(conn *sql.DB) ([]entities.Movie, error) {
	var movies []entities.Movie
	rows, err := conn.Query("SELECT id, name, release_date, director, description, duration, budget FROM movies")
	if err != nil {
		return nil, fmt.Errorf("GetAll: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var movie entities.Movie
		var description sql.NullString
		var duration, budget sql.NullInt32
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.ReleaseDate, &movie.Director, &description, &duration, &budget); err != nil {
			return nil, fmt.Errorf("GetAll: %v", err)
		}
		if description.Valid {
			movie.Description = description.String
		}
		if duration.Valid {
			movie.Duration = uint32(duration.Int32)
		}
		if budget.Valid {
			movie.Budget = uint32(budget.Int32)
		}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAll: %v", err)
	}
	return movies, nil
}

func GetById(conn *sql.DB, id uuid.UUID) (entities.Movie, error) {
	var movie entities.Movie
	var description sql.NullString
	var duration, budget sql.NullInt32
	row := conn.QueryRow("SELECT id, name, release_date, director, description, duration, budget FROM movies WHERE id = $1", id)
	if err := row.Scan(&movie.ID, &movie.Name, &movie.ReleaseDate, &movie.Director, &description, &duration, &budget); err != nil {
		if err == sql.ErrNoRows {
			return movie, fmt.Errorf("GetById %d: no such movie", id)
		}
		return movie, fmt.Errorf("GetById: %v", err)
	}
	if description.Valid {
		movie.Description = description.String
	}
	if duration.Valid {
		movie.Duration = uint32(duration.Int32)
	}
	if budget.Valid {
		movie.Budget = uint32(budget.Int32)
	}
	return movie, nil
}
