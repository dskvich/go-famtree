package domain

import (
	"time"
)

// swagger:model
type Tree struct {
	// swagger:strfmt uuid
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
