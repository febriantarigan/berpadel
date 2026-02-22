package response

import (
	"sort"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/util"
)

type TournamentResponse struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	Status    string          `json:"status"`
	Season    string          `json:"season"`
	Location  string          `json:"location"`
	Courts    []string        `json:"courts"`
	MaxPoints int             `json:"max_points"`
	Rounds    []RoundResponse `json:"rounds"`
	CreatedAt string          `json:"created_at"`
}

type RoundResponse struct {
	Round          int             `json:"round"`
	Matches        []MatchResponse `json:"matches"`
	WaitingPlayers []string        `json:"waiting_players"`
}

type MatchResponse struct {
	ID     string `json:"id"`
	Court  string `json:"court"`
	TeamA  *Team  `json:"team_a"`
	TeamB  *Team  `json:"team_b"`
	Status string `json:"status"`
}

type Team struct {
	Players []string `json:"players"`
	Score   int      `json:"score"`
}

func NewTournamentResponse(t *domain.Tournament, matches []*domain.Match) *TournamentResponse {
	return &TournamentResponse{
		ID:        t.ID,
		Name:      t.Name,
		Type:      string(t.Type),
		Status:    string(t.Status),
		Season:    t.Season,
		Location:  t.Location,
		Courts:    t.Courts,
		MaxPoints: t.MaxPoints,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
		Rounds:    mapRounds(matches),
	}
}

func mapRounds(matches []*domain.Match) []RoundResponse {
	roundsMap := make(map[int][]MatchResponse)

	for _, m := range matches {
		roundsMap[m.Round] = append(
			roundsMap[m.Round],
			mapMatchResponse(m),
		)
	}

	rounds := make([]RoundResponse, 0, len(roundsMap))
	for round, matches := range roundsMap {
		rounds = append(rounds, RoundResponse{
			Round:   round,
			Matches: matches,
		})
	}

	sort.Slice(rounds, func(i, j int) bool {
		return rounds[i].Round < rounds[j].Round
	})

	return rounds
}

func mapMatchResponse(m domain.Match) MatchResponse {
	return MatchResponse{
		ID:     m.ID,
		Court:  m.Court,
		TeamA:  mapTeamResponse(m.TeamA),
		TeamB:  mapTeamResponse(m.TeamB),
		Status: string(m.Status),
	}
}

func mapTeamResponse(t domain.Team) *Team {
	playerIDs := util.UsersToStrings(t.Players)
	return &Team{
		Players: playerIDs,
		Score:   t.Score,
	}
}
