package controller

import (
	"net/http"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"

	"github.com/google/uuid"

	"github.com/joffrua/go-famtree/internal/domain"
)

type TreeController struct {
}

func NewTreeController() *TreeController {
	return &TreeController{}
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
	uuid, _ := uuid.NewUUID()
	trees := []domain.Tree{
		{ID: uuid, Name: "My #1 tree", Description: "some description 1"},
		{ID: uuid, Name: "My #2 tree", Description: "some description 2"},
	}

	httpserver.RespondWithJSON(w, http.StatusOK, trees)
}

// GetTree godoc
// @Summary Get a tree
// @Description Get a tree by ID
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path int true "Tree ID"
// @Success 200 {object} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [get]
func (ctrl TreeController) GetTree(w http.ResponseWriter, r *http.Request) {
	uuid, _ := uuid.NewUUID()
	tree := domain.Tree{ID: uuid, Name: "My #1 tree", Description: "some description 1"}

	httpserver.RespondWithJSON(w, http.StatusOK, tree)
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
	uuid, _ := uuid.NewUUID()
	tree := domain.Tree{ID: uuid, Name: "My #1 tree", Description: "some description 1"}

	httpserver.RespondWithJSON(w, http.StatusCreated, tree)
}

// UpdateTree godoc
// @Summary Update a tree
// @Description Update by json tree
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path int true "Tree ID"
// @Param tree body domain.Tree true "Update tree"
// @Success 200 {object} domain.Tree
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [patch]
func (ctrl TreeController) UpdateTree(w http.ResponseWriter, r *http.Request) {
	uuid, _ := uuid.NewUUID()
	tree := domain.Tree{ID: uuid, Name: "My #1 tree", Description: "some description 1"}

	httpserver.RespondWithJSON(w, http.StatusOK, tree)
}

// DeleteTree godoc
// @Summary Delete a tree
// @Description Delete by tree ID
// @Tags Trees
// @Accept json
// @Produce json
// @Param tree_id path int true "Tree ID"
// @Success 204
// @Failure 400 {object} httpserver.Error
// @Failure 404 {object} httpserver.Error
// @Failure 500 {object} httpserver.Error
// @Router /trees/{tree_id} [delete]
func (ctrl TreeController) DeleteTree(w http.ResponseWriter, r *http.Request) {
	httpserver.RespondWithJSON(w, http.StatusNoContent, nil)
}
