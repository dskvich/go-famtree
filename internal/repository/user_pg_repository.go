package repository

import (
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

func (repo UserPgRepository) GetAll() ([]domain.User, error) {
	var u []domain.User
	err := repo.pg.GetConnection().Model(&u).Select()
	return u, err
}

func (repo UserPgRepository) GetByID(id string) (*domain.User, error) {
	var u domain.User
	err := repo.pg.GetConnection().Model(&u).Where("id = ?", id).Select()
	return &u, err
}

func (repo UserPgRepository) Persist(u *domain.User) error {
	_, err := repo.pg.GetConnection().Model(u).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (repo UserPgRepository) Delete(id string) error {
	_, err := repo.pg.GetConnection().Model((*domain.User)(nil)).Where("id = ?", id).Delete()
	return err
}
