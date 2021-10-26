package controller

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"

	"github.com/joffrua/go-famtree/internal/domain"
)

type UserController struct {
	repo domain.UserRepository
}

func NewUserController(repo domain.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

// swagger:route GET /users Users getAllUsers
//
// List all users
//
// List all users
//
// responses:
//   200: []User
//   400: Error
//   404: Error
//   500: Error
func (ctrl UserController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	u, err := ctrl.repo.FindAll()
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, u)
}

// swagger:route GET /users/{user_id} Users getUser
//
// Get a user
//
// Get a user by ID
//
// responses:
//   200: User
//   400: Error
//   404: Error
//   500: Error
func (ctrl UserController) GetUser(w http.ResponseWriter, _ *http.Request) {
	ID, _ := uuid.Parse("91cf6ac3-ec86-4e6f-8b60-7f3cff879ec5")

	u, err := ctrl.repo.FindByID(ID)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, u)
}
