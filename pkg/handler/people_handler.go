package handler

import (
	"context"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/people"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/sushkevichd/go-famtree/api/models"
	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type PeopleRepository interface {
	FindAllByTreeID(ctx context.Context, treeID uuid.UUID) ([]*domain.Person, error)
}

type PeopleHandler struct {
	repo PeopleRepository
}

func NewPeopleHandler(repo PeopleRepository) *PeopleHandler {
	return &PeopleHandler{
		repo: repo,
	}
}

func (h *PeopleHandler) GetPeople(params people.GetPeopleParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	treeID, err := uuid.Parse(params.TreeID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("parsing uuid")
		return people.NewGetPeopleDefault(403)
	}

	peopleList, err := h.repo.FindAllByTreeID(ctx, treeID)
	if err != nil {
		log.Err(err).Interface("params", params).Msg("finding all people")
		return people.NewGetPeopleDefault(500)
	}

	mappedPeople := make([]*models.Person, len(peopleList))
	for i, u := range peopleList {
		mappedPeople[i] = mapDomainPerson(u)
	}

	return people.NewGetPeopleOK().WithPayload(mappedPeople)
}

func mapDomainPerson(p *domain.Person) *models.Person {
	res := &models.Person{
		Name: &p.Name,
	}
	if p.ID != uuid.Nil {
		id := strfmt.UUID(p.ID.String())
		res.ID = id
	}
	if p.FatherID != uuid.Nil {
		fatherID := strfmt.UUID(p.FatherID.String())
		res.FatherID = fatherID
	}
	if p.MotherID != uuid.Nil {
		motherID := strfmt.UUID(p.MotherID.String())
		res.MotherID = motherID
	}
	return res
}
