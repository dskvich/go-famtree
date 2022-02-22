package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/joffrua/go-famtree/internal/domain"
	"github.com/joffrua/go-famtree/internal/infra/httpserver"
)

type UserController struct {
	repo domain.UserRepository
}

func NewUserController(repo domain.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (c UserController) New(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := httpserver.DecodeJSONBody(w, r, &user); err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	log.Infof("create user: %+v", user)

	if err := c.repo.Persist(&user); err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusCreated, user)
}

func (c UserController) GetAll(w http.ResponseWriter, _ *http.Request) {
	users, err := c.repo.FindAll()
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusOK, users)
}

func (c UserController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
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

func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := httpserver.DecodeJSONBody(w, r, &user); err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user.ID = ID

	err = c.repo.Persist(&user)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httpserver.RespondWithJSON(w, http.StatusOK, user)
}
