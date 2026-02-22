package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/oklog/ulid/v2"
)

// MixAmericanoGenerator creates matches for Mix Americano (rotating partners, mixed M-F).
// Each team must have one male and one female. Partners rotate each round.
// Requires 4, 8, 12, 16... players with equal M/F count.
type MixAmericanoGenerator struct {
	Courts    []string
	MaxPoints int
}

var (
	ErrMixAmericanoInvalidCount   = errors.New("mix_americano requires equal male/female count, divisible by 4 total (e.g. 4M+4F)")
	ErrMixAmericanoGenderImbalance = errors.New("mix_americano requires equal number of male and female players")
)

func (g *MixAmericanoGenerator) Generate(t *domain.Tournament, players []*domain.User) ([]*RoundSchedule, error) {
	males, females := splitByGender(players)
	if len(males) != len(females) {
		return nil, ErrMixAmericanoGenderImbalance
	}
	n := len(players)
	if n < 4 || n%4 != 0 {
		return nil, fmt.Errorf("%w: got %d players (need equal M/F)", ErrMixAmericanoInvalidCount, n)
	}

	shuffledMales := shufflePlayers(males)
	shuffledFemales := shufflePlayers(females)

	numRounds := n / 2 // Each male partners with each female exactly once
	schedules := make([]*RoundSchedule, 0, numRounds)

	for r := 0; r < numRounds; r++ {
		pairs := g.generateMixedPairsForRound(n/2, r, shuffledMales, shuffledFemales)
		teams := make([]*domain.Team, 0, n/2)
		for _, p := range pairs {
			teams = append(teams, &domain.Team{
				Players: []*domain.User{p.male, p.female},
			})
		}

		matches := []*domain.Match{}
		for i := 0; i < len(teams); i += 2 {
			court := g.Courts[(i/2)%len(g.Courts)]
			match := &domain.Match{
				ID:               ulid.Make().String(),
				TournamentID:     t.ID,
				Round:            r + 1,
				Court:            court,
				TeamA:            *teams[i],
				TeamB:            *teams[i+1],
				WaitingPlayerIDs: []string{},
				Status:           domain.MatchScheduled,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}
			matches = append(matches, match)
		}

		schedules = append(schedules, &RoundSchedule{
			Round:          r + 1,
			Matches:        matches,
			WaitingPlayers: []*domain.User{},
		})
	}

	return schedules, nil
}

type mixedPair struct{ male, female *domain.User }

func (g *MixAmericanoGenerator) generateMixedPairsForRound(nPairs, round int, males, females []*domain.User) []mixedPair {
	// Rotate females; fix male 0 with female round, then pair (male i, female (round+i) mod n)
	pairs := make([]mixedPair, 0, nPairs)
	for i := 0; i < nPairs; i++ {
		maleIdx := i
		femaleIdx := (round + i) % nPairs
		pairs = append(pairs, mixedPair{males[maleIdx], females[femaleIdx]})
	}
	return pairs
}

func splitByGender(players []*domain.User) (males, females []*domain.User) {
	for _, p := range players {
		if p.Gender == domain.GenderMale {
			males = append(males, p)
		} else {
			females = append(females, p)
		}
	}
	return males, females
}
