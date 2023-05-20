package database

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type PeopleBunRepository struct {
	db *bun.DB
}

func NewPeopleBunRepository(db *bun.DB) *PeopleBunRepository {
	return &PeopleBunRepository{
		db: db,
	}
}

func (r *PeopleBunRepository) FindPeopleByTreeID(ctx context.Context, treeID uuid.UUID) ([]*domain.Person, error) {
	var people []*domain.Person
	if err := r.db.NewSelect().Model(&people).Where("tree_id = ?", treeID).Scan(ctx); err != nil {
		return nil, err
	}
	return people, nil
}
