package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/oklog/ulid/v2"
)

// AmericanoFixedGenerator creates matches for Americano Fixed (fixed teams) tournaments.
// Teams are formed once at the start and stay the same throughout.
// Requires player count divisible by 4 (e.g. 8, 12, 16).
type AmericanoFixedGenerator struct {
	Courts    []string
	MaxPoints int
}

var ErrAmericanoFixedInvalidPlayerCount = errors.New("americano_fixed requires player count divisible by 4")

func (g *AmericanoFixedGenerator) Generate(t *domain.Tournament, players []*domain.User) ([]*RoundSchedule, error) {
	n := len(players)
	if n < 4 || n%4 != 0 {
		return nil, fmt.Errorf("%w: got %d players", ErrAmericanoFixedInvalidPlayerCount, n)
	}

	shuffled := shufflePlayers(players)
	teams := makeTeams(shuffled)

	courtSlots := len(g.Courts) * 2
	var schedules []*RoundSchedule
	round := 1

	for len(teams) > 0 {
		teamsInRound := teams[:min(len(teams), courtSlots)]
		teams = teams[len(teamsInRound):]

		waitingPlayerIDs := []string{}
		if len(teamsInRound)%2 != 0 {
			for _, p := range teamsInRound[len(teamsInRound)-1].Players {
				waitingPlayerIDs = append(waitingPlayerIDs, p.ID)
			}
		}

		matches := []*domain.Match{}
		for i := 0; i < len(teamsInRound)-1; i += 2 {
			court := g.Courts[(i/2)%len(g.Courts)]
			match := &domain.Match{
				ID:               ulid.Make().String(),
				TournamentID:     t.ID,
				Round:            round,
				Court:            court,
				TeamA:            *teamsInRound[i],
				TeamB:            *teamsInRound[i+1],
				WaitingPlayerIDs: waitingPlayerIDs,
				Status:           domain.MatchScheduled,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}
			matches = append(matches, match)
		}

		waitingPlayers := []*domain.User{}
		if len(teamsInRound)%2 != 0 {
			waitingPlayers = teamsInRound[len(teamsInRound)-1].Players
		}

		schedules = append(schedules, &RoundSchedule{
			Round:          round,
			Matches:        matches,
			WaitingPlayers: waitingPlayers,
		})
		round++
	}

	return schedules, nil
}
