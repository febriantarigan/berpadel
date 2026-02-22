package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/oklog/ulid/v2"
)

// AmericanoGenerator creates matches for Americano tournaments.
// In Americano format, each player partners with every other player exactly once
// across (n-1) rounds. Players rotate partners each round.
// Requires player count divisible by 4 (e.g. 8, 12, 16).
type AmericanoGenerator struct {
	Courts    []string
	MaxPoints int
}

// ErrInvalidPlayerCount is returned when player count is not valid for Americano.
var ErrInvalidPlayerCount = errors.New("americano requires player count divisible by 4 (e.g. 8, 12, 16)")

func (g *AmericanoGenerator) Generate(t *domain.Tournament, players []*domain.User) ([]*RoundSchedule, error) {
	n := len(players)
	if n < 4 || n%4 != 0 {
		return nil, fmt.Errorf("%w: got %d players", ErrInvalidPlayerCount, n)
	}

	// Shuffle once for fairness, then apply rotating partner algorithm
	shuffled := shufflePlayers(players)

	// Americano: n-1 rounds, each player partners with every other exactly once
	numRounds := n - 1
	schedules := make([]*RoundSchedule, 0, numRounds)

	for r := 0; r < numRounds; r++ {
		pairs := g.generatePairsForRound(n, r)
		teams := make([]*domain.Team, 0, n/2)
		for _, pair := range pairs {
			teams = append(teams, &domain.Team{
				Players: []*domain.User{shuffled[pair.a], shuffled[pair.b]},
			})
		}

		// Pair up teams into matches: (0,1) vs (2,3), (4,5) vs (6,7), etc.
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
				WaitingPlayerIDs: []string{}, // Americano: everyone plays each round
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

// pair represents two player indices who partner in a round
type pair struct{ a, b int }

// generatePairsForRound uses the standard round-robin doubles algorithm:
// Fix player n-1, pair with player r. Pair remaining (i+r+1) with (n-2-i+r+1) mod (n-1).
// This ensures each player partners with every other exactly once over n-1 rounds.
func (g *AmericanoGenerator) generatePairsForRound(n, round int) []pair {
	fixed := n - 1
	partner := round

	pairs := []pair{{fixed, partner}}

	used := make([]bool, n)
	used[fixed] = true
	used[partner] = true

	remaining := make([]int, 0, n-2)
	for i := 0; i < n; i++ {
		if !used[i] {
			remaining = append(remaining, i)
		}
	}

	// Pair consecutive in remaining (they form the other n/2 - 1 pairs)
	for i := 0; i < len(remaining); i += 2 {
		pairs = append(pairs, pair{remaining[i], remaining[i+1]})
	}

	return pairs
}

