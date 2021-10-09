package domain

import "github.com/google/uuid"

// User example
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserRepository interface {
	New(tree *User) error
	FindByID(id uuid.UUID) (*User, error)
	Update(id uuid.UUID, name, description string) error
	FindAll(userId uuid.UUID) ([]User, error)
	Remove(id uuid.UUID) error
}
