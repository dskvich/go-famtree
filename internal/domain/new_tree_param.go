package domain

// swagger:parameters newTree
type NewTreeParam struct {
	// required: true
	// in: body
	Tree Tree `json:"tree"`
}
