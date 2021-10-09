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

// GetAllTrees godoc
// @Summary List all trees
// @Description List all trees
// @Tags Trees
// @Accept json
// @Produce json
// @Success 200 {array} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees [get]
func (ctrl TreeController) GetAllTrees(w http.ResponseWriter, r *http.Request) {
	t, err := ctrl.repo.FindAll()
	if err != nil {
		httpserver.RespondWithError(w, http.StatusInternalServerError, err)
	}

	httpserver.RespondWithJSON(w, http.StatusOK, t)
}

// GetTree godoc
// @Summary Get a tree
// @Description Get a tree by ID
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path string true "Tree ID"
// @Success 200 {object} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [get]
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

// NewTree godoc
// @Summary Create a new tree
// @Description Create by json tree
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree body domain.Tree true "Create a new tree"
// @Success 201 {object} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees [post]
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

// UpdateTree godoc
// @Summary Update a tree
// @Description Update by json tree
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path string true "Tree ID"
// @Param tree body domain.Tree true "Update tree"
// @Success 200 {object} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [put]
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

// DeleteTree godoc
// @Summary Delete a tree
// @Description Delete by tree ID
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path string true "Tree ID"
// @Success 204
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [delete]
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
