package handler

// swagger:model
type errorResponse struct {
	Message string `json:"message,omitempty" example:"status bad request"`
}
