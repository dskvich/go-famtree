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

// GetAllUsers godoc
// @Summary List all users
// @Description List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /users [get]
func (ctrl UserController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	u, err := ctrl.repo.FindAll()
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, u)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /users/{user_id} [get]
func (ctrl UserController) GetUser(w http.ResponseWriter, _ *http.Request) {
	ID, _ := uuid.Parse("91cf6ac3-ec86-4e6f-8b60-7f3cff879ec5")

	u, err := ctrl.repo.FindByID(ID)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, u)
}
