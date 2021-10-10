package domain

import (
	"time"

	"github.com/google/uuid"
)

// Person example
type Person struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	TreeID     uuid.UUID `json:"tree_id"`
	FatherID   uuid.UUID `json:"father_id"`
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
