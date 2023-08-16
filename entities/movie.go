package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

type Movie struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	Director    string    `json:"director"`
	Description string    `json:"description,omitempty"`
	Duration    uint      `json:"duration,omitempty"`
	Budget      uint      `json:"budget,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
