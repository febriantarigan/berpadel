package repository

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type MatchRepository interface {
	BatchCreate(ctx context.Context, matches []*domain.Match) error
	GetByID(ctx context.Context, id string) (*domain.Match, error)
	SubmitScore(ctx context.Context, match domain.Match) error
	ListByTournamentID(ctx context.Context, tournamentID string) ([]*domain.Match, error)
}
