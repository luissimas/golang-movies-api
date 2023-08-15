package entities

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	Director    string    `json:"director"`
	Description string    `json:"description,omitempty"`
	Duration    uint32    `json:"duration,omitempty"`
	Budget      uint32    `json:"budget,omitempty"`
}
