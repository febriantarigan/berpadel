package repository

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByIDs(ctx context.Context, ids []string) ([]*domain.User, error)
	SearchByName(ctx context.Context, keyword string) ([]*domain.User, error)
	List(ctx context.Context) ([]*domain.User, error)
}
