package handler

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/go-openapi/strfmt"

	"github.com/sushkevichd/go-famtree/api/models"

	"github.com/go-openapi/runtime/middleware"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/users"

	"github.com/google/uuid"

	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type UserRepository interface {
	FindAll(context.Context) ([]domain.User, error)
	FindByID(context.Context, uuid.UUID) (*domain.User, error)
	Persist(context.Context, *domain.User) error
	Delete(context.Context, uuid.UUID) error
}

type UserHandler struct {
	repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (c *UserHandler) CreateUser(params users.CreateUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	user := mapModelUser(params.User)

	if err := c.repo.Persist(ctx, user); err != nil {
		log.Err(err).Interface("params", params).Msg("creating a user")
		return users.NewCreateUserDefault(500)
	}

	return users.NewCreateUserCreated().WithPayload(mapDomainUser(*user))
}

func (c *UserHandler) GetUsers(params users.GetUsersParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	userList, err := c.repo.FindAll(ctx)
	if err != nil {
		log.Err(err).Interface("params", params).Msg("finding all users")
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

func (c *UserHandler) GetUserByID(params users.GetUserByIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("parsing uuid")
		return users.NewGetUserByIDDefault(403)
	}

	res, err := c.repo.FindByID(ctx, id)
	if err != nil {
		log.Err(err).Interface("params", params).Msg("finding a user by id")
		if err == sql.ErrNoRows {
			return users.NewGetUserByIDDefault(404)
		}
		return users.NewGetUserByIDDefault(500)
	}

	return users.NewGetUserByIDOK().WithPayload(mapDomainUser(*res))
}

func (c *UserHandler) DeleteUserByID(params users.DeleteUserByIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("deleting a user by id")
		return users.NewDeleteUserByIDDefault(403)
	}

	if err = c.repo.Delete(ctx, id); err != nil {
		log.Err(err).Interface("params", params).Msg("deleting a user by id")
		if err == sql.ErrNoRows {
			return users.NewDeleteUserByIDDefault(404)
		}
		return users.NewDeleteUserByIDDefault(500)
	}

	return users.NewDeleteUserByIDNoContent()
}

func (c *UserHandler) UpdateUserByID(params users.UpdateUserByIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	id, err := uuid.Parse(params.UserID.String())
	if err != nil {
		log.Err(err).Interface("params", params).Msg("updating a user by id")
		return users.NewUpdateUserByIDDefault(403)
	}

	user := mapModelUser(params.User)
	user.ID = &id

	if err := c.repo.Persist(ctx, user); err != nil {
		log.Err(err).Interface("params", params).Msg("updating a user by id")
		return users.NewUpdateUserByIDDefault(500)
	}

	return users.NewUpdateUserByIDNoContent()
}
