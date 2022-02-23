package db

import (
	"context"

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

func (repo TreePgRepository) FindAll() (t []domain.Tree, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&t).Scan(context.Background())
	return
}

func (repo TreePgRepository) FindByID(id uuid.UUID) (t *domain.Tree, err error) {
	err = repo.pg.GetConnection().NewSelect().Model(&t).Where("id = ?", id).Scan(context.Background())
	return
}

func (repo TreePgRepository) Persist(t *domain.Tree) error {
	_, err := repo.pg.GetConnection().NewInsert().Model(t).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	return err
}

func (repo TreePgRepository) Delete(id uuid.UUID) error {
	_, err := repo.pg.GetConnection().NewDelete().Model((*domain.Tree)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}
