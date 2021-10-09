package domain

import "github.com/google/uuid"

// Tree example
type Tree struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type TreeRepository interface {
	New(tree *Tree) error
	FindByID(id uuid.UUID) (*Tree, error)
	Update(id uuid.UUID, name, description string) error
	FindAll(userId uuid.UUID) ([]Tree, error)
	Remove(id uuid.UUID) error
}
