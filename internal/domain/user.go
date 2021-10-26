package domain

import (
	"time"

	"github.com/google/uuid"
)

// User example
type User struct {
	ID           uuid.UUID `json:"id" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	Login        string    `json:"login"`
	Role         string    `json:"role"`
	Lang         string    `json:"lang"`
	Name         string    `json:"name"`
	Email        string    `json:"email" format:"email"`
	PasswordHash string    `json:"-"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(uuid.UUID) (*User, error)
	Persist(*User) error
	Delete(uuid.UUID) error
}
