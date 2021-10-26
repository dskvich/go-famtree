package db

import (
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

func (repo UserPgRepository) FindAll() ([]domain.User, error) {
	u := []domain.User{}
	err := repo.pg.GetConnection().Model(&u).Select()
	return u, err
}

func (repo UserPgRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	u := domain.User{}
	err := repo.pg.GetConnection().Model(&u).Where("id = ?", id).Select()
	return &u, err
}

func (repo UserPgRepository) Persist(u *domain.User) error {
	_, err := repo.pg.GetConnection().Model(u).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (repo UserPgRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().Model((*domain.User)(nil)).Where("id = ?", id).Delete()
	return err
}
