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
	t := []domain.Tree{}
	err := repo.pg.GetConnection().Model(&t).Select()
	return t, err
}

func (repo TreePgRepository) FindByID(id uuid.UUID) (*domain.Tree, error) {
	t := domain.Tree{}
	err := repo.pg.GetConnection().Model(&t).Where("id = ?", id).Select()
	return &t, err
}

func (repo TreePgRepository) Persist(t *domain.Tree) error {
	_, err := repo.pg.GetConnection().Model(t).OnConflict("(id) DO UPDATE").Insert()
	return err
}

func (repo TreePgRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().Model((*domain.Tree)(nil)).Where("id = ?", id).Delete()
	return err
}
