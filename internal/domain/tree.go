package domain

import "github.com/google/uuid"

// Tree example
type Tree struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type TreeRepository interface {
	FindAll() ([]Tree, error)
	FindByID(uuid.UUID) (*Tree, error)
	Persist(*Tree) error
	Delete(uuid.UUID) error
}
