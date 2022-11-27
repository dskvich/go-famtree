package repository

import (
	"github.com/joffrua/go-famtree/internal/domain"
)

type TreePgRepository struct {
	pg *Pg
}

func NewTreePgRepository(pg *Pg) *TreePgRepository {
	return &TreePgRepository{
		pg: pg,
	}
}

func (repo TreePgRepository) GetAll() ([]domain.Tree, error) {
	var u []domain.Tree
	err := repo.pg.GetConnection().Model(&u).Select()
	return u, err
}

func (repo TreePgRepository) GetByID(id string) (*domain.Tree, error) {
	var u domain.Tree
	err := repo.pg.GetConnection().Model(&u).Where("id = ?", id).Select()
	return &u, err
}

func (repo TreePgRepository) Persist(u *domain.Tree) error {
	_, err := repo.pg.GetConnection().Model(u).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (repo TreePgRepository) Delete(id string) error {
	_, err := repo.pg.GetConnection().Model((*domain.Tree)(nil)).Where("id = ?", id).Delete()
	return err
}
