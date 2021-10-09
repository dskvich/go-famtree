package domain

import "github.com/google/uuid"

// User example
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(uuid.UUID) (*User, error)
	Persist(*User) error
	Delete(uuid.UUID) error
}
