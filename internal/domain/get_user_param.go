package domain

import "github.com/google/uuid"

// swagger:parameters getUser
type GetUserParam struct {
	// required: true
	// in: path
	// swagger:strfmt uuid
	UserID uuid.UUID `json:"user_id"`
}
