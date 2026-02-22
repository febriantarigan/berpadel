package handler

import (
	"github.com/febriantarigan/berpadel/internal/service"
)

type LeaderboardHandler struct {
	leaderboardService *service.LeaderboardService
}

func NewLeaderboardHandler(lb *service.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{leaderboardService: lb}
}
