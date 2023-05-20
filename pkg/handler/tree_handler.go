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

func (h *TreeHandler) GetAllTreesForUser(params trees.GetAllTreesForUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userID, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("parsing uuid")
		return trees.NewGetAllTreesForUserDefault(403)
	}

	treeList, err := h.repo.FindAllByUserID(ctx, userID)
	if err != nil {
		log.Err(err).Interface("params", params).Msg("finding all users")
		return trees.NewGetAllTreesForUserDefault(500)
	}

	mappedTrees := make([]*models.Tree, len(treeList))
	for i, u := range treeList {
		mappedTrees[i] = mapDomainTree(u)
	}

	return trees.NewGetAllTreesForUserOK().WithPayload(mappedTrees)
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

func (h *TreeHandler) CreateTreeForUser(params trees.CreateTreeForUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	userID, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("parsing uuid")
		return trees.NewCreateTreeForUserDefault(403)
	}
	tree := mapModelTree(params.Tree)
	tree.UserID = &userID

	if err := h.repo.Persist(ctx, tree); err != nil {
		log.Err(err).Interface("params", params).Msg("creating a tree for user")
		return trees.NewCreateTreeForUserDefault(500)
	}

	return trees.NewCreateTreeForUserCreated().WithPayload(mapDomainTree(tree))
}
