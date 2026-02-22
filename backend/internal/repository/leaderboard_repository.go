package repository

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type LeaderboardRepository interface {
	Create(ctx context.Context, lb domain.Leaderboard) error
	Update(ctx context.Context, lb domain.Leaderboard) error
	GetByTournamentID(ctx context.Context, tournamentID string) (*domain.Leaderboard, error)
	List(ctx context.Context) ([]domain.Leaderboard, error)
}
