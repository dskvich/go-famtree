package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"

	"github.com/google/uuid"

	"github.com/joffrua/go-famtree/internal/domain"
)

type TreeController struct {
	repo domain.TreeRepository
}

func NewTreeController(repo domain.TreeRepository) *TreeController {
	return &TreeController{
		repo: repo,
	}
}

// swagger:route GET /trees Trees getAllTrees
//
// List all trees
//
// List all trees
//
// responses:
//   200: []Tree
//   400: Error
//   404: Error
//   500: Error
func (ctrl TreeController) GetAllTrees(w http.ResponseWriter, _ *http.Request) {
	t, err := ctrl.repo.FindAll()
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, t)
}

// swagger:route GET /trees/{tree_id} Trees getTree
//
// Get a tree
//
// Get a tree by ID
//
// responses:
//   200: Tree
//   400: Error
//   404: Error
//   500: Error
func (ctrl TreeController) GetTree(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
	}

	t, err := ctrl.repo.FindByID(ID)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, t)
}

// swagger:route POST /trees Trees newTree
//
// Create a new tree
//
// Create by json tree
//
// responses:
//   201: Tree
//   400: Error
//   404: Error
//   500: Error
func (ctrl TreeController) NewTree(w http.ResponseWriter, r *http.Request) {
	var t domain.Tree

	err := httpserver.DecodeJSONBody(w, r, &t)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
	}

	err = ctrl.repo.Persist(&t)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusCreated, t)
}

// swagger:route PUT /trees/{tree_id} Trees updateTree
//
// Update a tree
//
// Update by json tree
//
// responses:
//   200: Tree
//   400: Error
//   404: Error
//   500: Error
func (ctrl TreeController) UpdateTree(w http.ResponseWriter, r *http.Request) {
	var t domain.Tree

	err := httpserver.DecodeJSONBody(w, r, &t)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
	}

	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
	}

	t.ID = ID

	err = ctrl.repo.Persist(&t)
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, t)
}

// swagger:route DELETE /trees/{tree_id} Trees deleteTree
//
// Delete a tree
//
// Delete by tree ID
//
// responses:
//   204:
//   400: Error
//   404: Error
//   500: Error
func (ctrl TreeController) DeleteTree(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		httpserver.RespondWithError(w, http.StatusBadRequest, err)
	}

	if err = ctrl.repo.Delete(ID); err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusNoContent, nil)
}
