package service

import (
	"math/rand"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
)

func shufflePlayers(players []*domain.User) []*domain.User {
	shuffled := make([]*domain.User, len(players))
	copy(shuffled, players)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	return shuffled
}

func makeTeams(players []*domain.User) []*domain.Team {
	teams := make([]*domain.Team, 0, len(players)/2)
	for i := 0; i < len(players)-1; i += 2 {
		teams = append(teams, &domain.Team{
			Players: []*domain.User{players[i], players[i+1]},
		})
	}
	return teams
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
