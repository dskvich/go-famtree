package domain

import "github.com/google/uuid"

// swagger:parameters getTree
type GetTreeParam struct {
	// required: true
	// in: path
	// swagger:strfmt uuid
	TreeID uuid.UUID `json:"tree_id"`
}
