package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"

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

func (c UserHandler) CreateUser(param users.CreateUserParams) middleware.Responder {
	user := mapModelUser(param.User)

	if err := c.repo.Persist(user); err != nil {
		log.Err(err).Interface("param", param).Msg("user creation")
		return users.NewCreateUserDefault(500)
	}

	return users.NewCreateUserCreated().WithPayload(mapDomainUser(*user))
}

func (c UserHandler) GetUsers(param users.GetUsersParams) middleware.Responder {
	ctx := param.HTTPRequest.Context()

	userList, err := c.repo.FindAll(ctx)
	if err != nil {
		return users.NewGetUsersDefault(500)
	}

	res := make([]*models.User, len(userList))
	for i, u := range userList {
		res[i] = mapDomainUser(u)
	}

	return users.NewGetUsersOK().WithPayload(res)
}

func mapDomainUser(u domain.User) *models.User {
	id := strfmt.UUID(u.ID.String())
	return &models.User{
		ID:    id,
		Login: &u.Login,
		Name:  &u.Name,
	}
}

func mapModelUser(u *models.User) *domain.User {
	return &domain.User{
		Login: *u.Login,
		Name:  *u.Name,
	}
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
