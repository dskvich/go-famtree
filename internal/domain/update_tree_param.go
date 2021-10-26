package domain

import "github.com/google/uuid"

// swagger:parameters updateTree
type UpdateTreeParam struct {
	// required: true
	// in: path
	// swagger:strfmt uuid
	TreeID uuid.UUID `json:"tree_id"`
	// required: true
	// in: body
	Tree Tree `json:"tree"`
}
