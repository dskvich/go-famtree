package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type UserBunRepository struct {
	pg *Pg
}

func NewUserBunRepository(pg *Pg) *UserBunRepository {
	return &UserBunRepository{
		pg: pg,
	}
}

func (repo UserBunRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	users := make([]domain.User, 0)
	if err := repo.pg.GetConnection().NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo UserBunRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := new(domain.User)
	if err := repo.pg.GetConnection().NewSelect().Model(user).Where("id = ?", id.String()).Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo UserBunRepository) Persist(ctx context.Context, u *domain.User) error {
	_, err := repo.pg.GetConnection().NewInsert().Model(u).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	return err
}

func (repo UserBunRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.GetConnection().NewDelete().Model((*domain.User)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
