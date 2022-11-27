package domain

import (
	"time"
)

// swagger:model
type User struct {
	ID        string    `json:"id" format:"uuid"`
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
