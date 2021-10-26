package domain

import (
	"time"

	"github.com/google/uuid"
)

// Person example
type Person struct {
	ID         uuid.UUID `json:"id" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	TreeID     uuid.UUID `json:"tree_id" format:"uuid"`
	FatherID   uuid.UUID `json:"father_id" format:"uuid"`
	MotherID   uuid.UUID `json:"mother_id" format:"uuid"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Patronymic string    `json:"patronymic"`
	MaidenName string    `json:"maiden_name"`
	BirthDate  time.Time `json:"birth_date" format:"date"`
	DeathDate  time.Time `json:"death_date" format:"date"`
	Bio        string    `json:"bio"`
}

type PersonRepository interface {
	FindAll() ([]Person, error)
	FindByID(uuid.UUID) (*Person, error)
	Persist(*Person) error
	Delete(uuid.UUID) error
}
