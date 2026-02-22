package repository

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type TournamentRepository interface {
	Create(ctx context.Context, t *domain.Tournament) error
	CreateWithNewUsers(ctx context.Context, t *domain.Tournament, users []*domain.User) error
	GetByID(ctx context.Context, tournamentID string) (*domain.Tournament, error)
	List(ctx context.Context) ([]*domain.Tournament, error)
	Update(ctx context.Context, tournamentID string, upd TournamentUpdate) error
}

type TournamentUpdate struct {
	Status   *domain.TournamentStatus
	IsActive *bool
}
