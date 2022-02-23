package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/joffrua/go-famtree/internal/domain"
)

type UserPgRepository struct {
	pg *Pg
}

func NewUserPgRepository(pg *Pg) *UserPgRepository {
	return &UserPgRepository{
		pg: pg,
	}
}

func (repo UserPgRepository) FindAll() (u []domain.User, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&u).Scan(context.Background())
	return
}

func (repo UserPgRepository) FindByID(id uuid.UUID) (u *domain.User, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&u).Where("id = ?", id).Scan(context.Background())
	return
}

func (repo UserPgRepository) Persist(u *domain.User) error {
	_, err := repo.pg.GetConnection().NewInsert().Model(u).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	return err
}

func (repo UserPgRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().NewDelete().Model((*domain.User)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}
