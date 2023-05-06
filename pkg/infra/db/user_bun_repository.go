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

func (repo UserBunRepository) FindAll(ctx context.Context) (u []domain.User, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&u).Scan(ctx)
	return
}

func (repo UserBunRepository) FindByID(id uuid.UUID) (u *domain.User, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&u).Where("id = ?", id).Scan(context.Background())
	return
}

func (repo UserBunRepository) Persist(u *domain.User) error {
	_, err := repo.pg.GetConnection().NewInsert().Model(u).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	return err
}

func (repo UserBunRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().NewDelete().Model((*domain.User)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}
