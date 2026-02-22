package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/oklog/ulid/v2"
)

// MixAmericanoFixedGenerator creates matches for Mix Americano Fixed (fixed teams, mixed M-F).
// Each team is one male + one female, formed once and fixed for the tournament.
// Requires 4, 8, 12, 16... players with equal M/F count.
type MixAmericanoFixedGenerator struct {
	Courts    []string
	MaxPoints int
}

var (
	ErrMixAmericanoFixedInvalidCount   = errors.New("mix_americano_fixed requires equal male/female count, divisible by 4 total")
	ErrMixAmericanoFixedGenderImbalance = errors.New("mix_americano_fixed requires equal number of male and female players")
)

func (g *MixAmericanoFixedGenerator) Generate(t *domain.Tournament, players []*domain.User) ([]*RoundSchedule, error) {
	males, females := splitByGender(players)
	if len(males) != len(females) {
		return nil, ErrMixAmericanoFixedGenderImbalance
	}
	n := len(players)
	if n < 4 || n%4 != 0 {
		return nil, fmt.Errorf("%w: got %d players", ErrMixAmericanoFixedInvalidCount, n)
	}

	shuffledMales := shufflePlayers(males)
	shuffledFemales := shufflePlayers(females)

	teams := make([]*domain.Team, 0, n/2)
	for i := 0; i < len(shuffledMales); i++ {
		teams = append(teams, &domain.Team{
			Players: []*domain.User{shuffledMales[i], shuffledFemales[i]},
		})
	}

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
