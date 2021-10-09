package db

import (
	"github.com/google/uuid"
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

func (repo TreePgRepository) FindAll() ([]domain.Tree, error) {
	var u []domain.Tree
	err := repo.pg.GetConnection().Model(&u).Select()
	return u, err
}

func (repo TreePgRepository) FindByID(ID uuid.UUID) (*domain.Tree, error) {
	var u domain.Tree
	err := repo.pg.GetConnection().Model(&u).Where("id = ?", ID).Select()
	return &u, err
}

func (repo TreePgRepository) Persist(u *domain.Tree) error {
	_, err := repo.pg.GetConnection().Model(u).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (repo TreePgRepository) Delete(ID uuid.UUID) error {
	_, err := repo.pg.GetConnection().Model((*domain.Tree)(nil)).Where("id = ?", ID).Delete()
	return err
}
