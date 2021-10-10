package domain

import (
	"time"

	"github.com/google/uuid"
)

// Tree example
type Tree struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type TreeRepository interface {
	FindAll() ([]Tree, error)
	FindByID(uuid.UUID) (*Tree, error)
	Persist(*Tree) error
	Delete(uuid.UUID) error
}
