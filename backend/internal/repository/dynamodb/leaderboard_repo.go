package dynamodb

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
)

type LeaderboardRepository struct {
	*BaseRepository
}

func NewLeaderboardRepository(base *BaseRepository) *LeaderboardRepository {
	return &LeaderboardRepository{BaseRepository: base}
}

func (*LeaderboardRepository) Create(ctx context.Context, lb domain.Leaderboard) error {
	return nil
}

func (*LeaderboardRepository) Update(ctx context.Context, lb domain.Leaderboard) error {
	return nil
}

func (*LeaderboardRepository) GetByTournamentID(ctx context.Context, tournamentID string) (*domain.Leaderboard, error) {
	return nil, nil
}

func (*LeaderboardRepository) List(ctx context.Context) ([]domain.Leaderboard, error) {
	return nil, nil
}
