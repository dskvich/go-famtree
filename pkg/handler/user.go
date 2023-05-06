package handler

import (
	"net/http"

	"github.com/go-openapi/strfmt"

	"github.com/sushkevichd/go-famtree/api/models"

	"github.com/go-openapi/runtime/middleware"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/users"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sushkevichd/go-famtree/pkg/domain"
	"github.com/sushkevichd/go-famtree/pkg/infra/httpserver"
)

type UserHandler struct {
	repo domain.UserRepository
}

func NewUserHandler(repo domain.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (c UserHandler) New(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := httpserver.DecodeJSONBody(w, r, &user); err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.repo.Persist(&user); err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusCreated, user)
}

func (c UserHandler) ListUsers(param users.ListUsersParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()
	userList, err := c.repo.FindAll(ctx)
	if err != nil {
		return users.NewListUsersDefault(500)
	}

	return users.NewListUsersOK().WithPayload(toUserModels(userList))
}

func toUserModels(users []domain.User) models.Users {
	res := make(models.Users, len(users))
	for i, u := range users {
		id := strfmt.UUID(u.ID.String())
		login := u.Login
		name := u.Name
		res[i] = &models.User{
			ID:    &id,
			Login: &login,
			Name:  &name,
		}
	}
	return res
}

func (c UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.repo.FindByID(ID)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusOK, user)
}

func (c UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err = c.repo.Delete(ID); err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusNoContent, nil)
}

func (c UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	var user domain.User
	if err := httpserver.DecodeJSONBody(w, r, &user); err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user.ID = ID

	if err := c.repo.Persist(&user); err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusOK, user)
}
