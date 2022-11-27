package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joffrua/go-famtree/internal/infrastructure/httpserver"

	"github.com/joffrua/go-famtree/internal/domain"
)

type TreeRepository interface {
	GetAll() ([]domain.Tree, error)
	GetByID(string) (*domain.Tree, error)
	Persist(*domain.Tree) error
	Delete(string) error
}

type TreeHandler struct {
	repo TreeRepository
}

func NewTreeHandler(repo TreeRepository) *TreeHandler {
	return &TreeHandler{
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
func (t TreeHandler) GetAllTrees(w http.ResponseWriter, _ *http.Request) {
	trees, err := t.repo.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, trees)
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
func (t TreeHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	ID, found := mux.Vars(r)["id"]
	if !found {
		respondWithError(w, http.StatusBadRequest, nil)
		return
	}

	tree, err := t.repo.GetByID(ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, tree)
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
func (t TreeHandler) NewTree(w http.ResponseWriter, r *http.Request) {
	var tree domain.Tree

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	if err := json.NewDecoder(r.Body).Decode(&tree); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if err := t.repo.Persist(&tree); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 200)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, t)
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
func (t TreeHandler) UpdateTree(w http.ResponseWriter, r *http.Request) {
	var tree domain.Tree
	if err := httpserver.DecodeJSONBody(w, r, &tree); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	ID, found := mux.Vars(r)["id"]
	if !found {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("path parameter id is missing"))
		return
	}

	tree.ID = ID

	if err := t.repo.Persist(&tree); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, t)
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
func (t TreeHandler) DeleteTree(w http.ResponseWriter, r *http.Request) {
	ID, found := mux.Vars(r)["id"]
	if !found {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("path parameter id is missing"))
		return
	}

	if err := t.repo.Delete(ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
