package dynamodb

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(base *BaseRepository) *UserRepository {
	return &UserRepository{BaseRepository: base}
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) CreateBatch(ctx context.Context, users []*domain.User) error {
	return nil
}

func (u *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) GetByIDs(ctx context.Context, ids []string) ([]*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) SearchByName(ctx context.Context, keyword string) ([]*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) List(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
