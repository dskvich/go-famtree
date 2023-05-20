package domain

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Name      string    `json:"name"`
	TreeID    uuid.UUID `json:"tree_id" format:"uuid"`
	FatherID  uuid.UUID `json:"father_id" format:"uuid"`
	MotherID  uuid.UUID `json:"mother_id" format:"uuid"`
}
