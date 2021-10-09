package controller

import (
	"net/http"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"

	"github.com/joffrua/go-famtree/internal/domain"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
func (ctrl UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Mary", LastName: "Jane"},
	}

	httpserver.RespondWithJSON(w, http.StatusOK, users)
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
func (ctrl UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	user := domain.User{FirstName: "John", LastName: "Doe"}

	httpserver.RespondWithJSON(w, http.StatusOK, user)
}
