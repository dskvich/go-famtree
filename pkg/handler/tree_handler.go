package handler

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/sushkevichd/go-famtree/api/models"
	"github.com/sushkevichd/go-famtree/api/restapi/operations/trees"
	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type TreeRepository interface {
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Tree, error)
	Persist(context.Context, *domain.Tree) error
	Delete(context.Context, uuid.UUID) error
}

type TreeHandler struct {
	repo TreeRepository
}

func NewTreeHandler(repo TreeRepository) *TreeHandler {
	return &TreeHandler{
		repo: repo,
	}
}

func (h *TreeHandler) GetTrees(params trees.GetTreesParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userID, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("parsing uuid")
		return trees.NewGetTreesDefault(403)
	}

	treeList, err := h.repo.FindAllByUserID(ctx, userID)
	if err != nil {
		log.Err(err).Interface("params", params).Msg("finding all trees")
		return trees.NewGetTreesDefault(500)
	}

	mappedTrees := make([]*models.Tree, len(treeList))
	for i, u := range treeList {
		mappedTrees[i] = mapDomainTree(u)
	}

	return trees.NewGetTreesOK().WithPayload(mappedTrees)
}

func mapDomainTree(t *domain.Tree) *models.Tree {
	id := strfmt.UUID(t.ID.String())
	userID := strfmt.UUID(t.UserID.String())
	return &models.Tree{
		ID:     id,
		Name:   &t.Name,
		UserID: &userID,
	}
}

func mapModelTree(t *models.Tree) *domain.Tree {
	return &domain.Tree{
		Name: *t.Name,
	}
}

func (h *TreeHandler) CreateTree(params trees.CreateTreeParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	tree := mapModelTree(params.Tree)

	if err := h.repo.Persist(ctx, tree); err != nil {
		log.Err(err).Interface("params", params).Msg("creating a tree")
		return trees.NewCreateTreeDefault(500)
	}

	return trees.NewCreateTreeCreated().WithPayload(mapDomainTree(tree))
}
