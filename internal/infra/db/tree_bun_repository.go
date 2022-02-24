package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/joffrua/go-famtree/internal/domain"
)

type TreeBunRepository struct {
	pg *Pg
}

func NewTreeBunRepository(pg *Pg) *TreeBunRepository {
	return &TreeBunRepository{
		pg: pg,
	}
}

func (repo TreeBunRepository) FindAll() (t []domain.Tree, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&t).Scan(context.Background())
	return
}

func (repo TreeBunRepository) FindByID(id uuid.UUID) (t *domain.Tree, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&t).Where("id = ?", id).Scan(context.Background())
	return
}

func (repo TreeBunRepository) Persist(t *domain.Tree) error {
	_, err := repo.pg.GetConnection().NewInsert().Model(t).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	return err
}

func (repo TreeBunRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().NewDelete().Model((*domain.Tree)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}
