package service

import (
	"github.com/febriantarigan/berpadel/internal/repository"
)

type LeaderboardService struct {
	leaderboardRepo repository.LeaderboardRepository
}

func NewLeaderboardService(leaderboardRepo repository.LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{
		leaderboardRepo: leaderboardRepo,
	}
}
