package database

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type TreeBunRepository struct {
	db *bun.DB
}

func NewTreeBunRepository(db *bun.DB) *TreeBunRepository {
	return &TreeBunRepository{
		db: db,
	}
}

func (r *TreeBunRepository) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Tree, error) {
	var trees []*domain.Tree
	if err := r.db.NewSelect().Model(&trees).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}
	return trees, nil
}

func (r *TreeBunRepository) Persist(ctx context.Context, t *domain.Tree) error {
	_, err := r.db.NewInsert().Model(t).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	return err
}

func (r *TreeBunRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.NewDelete().Model((*domain.Tree)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
