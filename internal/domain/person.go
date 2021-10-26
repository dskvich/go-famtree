package domain

import (
	"time"

	"github.com/google/uuid"
)

// swagger:model
type Person struct {
	// swagger:strfmt uuid
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// swagger:strfmt uuid
	TreeID uuid.UUID `json:"tree_id"`
	// swagger:strfmt uuid
	FatherID uuid.UUID `json:"father_id"`
	// swagger:strfmt uuid
	MotherID   uuid.UUID `json:"mother_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Patronymic string    `json:"patronymic"`
	MaidenName string    `json:"maiden_name"`
	BirthDate  string    `json:"birth_date"`
	DeathDate  string    `json:"death_date"`
	Bio        string    `json:"bio"`
}

type PersonRepository interface {
	FindAll() ([]Person, error)
	FindByID(uuid.UUID) (*Person, error)
	Persist(*Person) error
	Delete(uuid.UUID) error
}
