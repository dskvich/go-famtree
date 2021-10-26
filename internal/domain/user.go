package domain

import (
	"time"

	"github.com/google/uuid"
)

// swagger:model
type User struct {
	ID        uuid.UUID `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	Login     string    `json:"login"`
	Role      string    `json:"role"`
	Lang      string    `json:"lang"`
	Name      string    `json:"name"`
	// swagger:strfmt email
	Email string `json:"email"`
	// swagger:ignore
	PasswordHash string `json:"password_hash"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(uuid.UUID) (*User, error)
	Persist(*User) error
	Delete(uuid.UUID) error
}
