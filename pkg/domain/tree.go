package domain

import (
	"time"

	"github.com/google/uuid"
)

type Tree struct {
	ID          *uuid.UUID `json:"id" format:"uuid" bun:",pk"`
	CreatedAt   *time.Time `json:"created_at" format:"date-time"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UserID      *uuid.UUID `json:"user_id" format:"uuid"`
	RootID      *uuid.UUID `json:"root_id" format:"uuid"`
}
