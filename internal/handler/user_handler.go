package handler

import (
	"net/http"

	"github.com/joffrua/go-famtree/internal/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(string) (*domain.User, error)
	Persist(*domain.User) error
	Delete(string) error
}

type UserHandler struct {
	repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
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
func (u *UserHandler) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := u.repo.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
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
func (u *UserHandler) GetUser(w http.ResponseWriter, _ *http.Request) {
	user, err := u.repo.GetByID("91cf6ac3-ec86-4e6f-8b60-7f3cff879ec5")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
