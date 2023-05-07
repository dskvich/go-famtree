package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/sushkevichd/go-famtree/pkg/domain"
)

type UserBunRepository struct {
	db *bun.DB
}

func NewUserBunRepository(db *bun.DB) *UserBunRepository {
	return &UserBunRepository{
		db: db,
	}
}

func (r *UserBunRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserBunRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user := new(domain.User)
	if err := r.db.NewSelect().Model(user).Where("id = ?", id.String()).Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserBunRepository) Persist(ctx context.Context, user *domain.User) error {
	_, err := r.db.NewInsert().Model(user).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	return err
}

func (r *UserBunRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.NewDelete().Model((*domain.User)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}
