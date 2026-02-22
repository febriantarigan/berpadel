package service

import "github.com/febriantarigan/berpadel/internal/domain"

// RoundSchedule represents a round with matches and waiting teams
type RoundSchedule struct {
	Round          int
	Matches        []*domain.Match
	WaitingPlayers []*domain.User
}

type MatchGenerator interface {
	Generate(tournament *domain.Tournament, players []*domain.User) ([]*RoundSchedule, error)
}

// Factory to get generator by tournament type
func GetGenerator(tType domain.TournamentType, courts []string, maxPoints int) MatchGenerator {
	switch tType {
	case domain.TournamentTypeAmericano:
		return &AmericanoGenerator{Courts: courts, MaxPoints: maxPoints}
	case domain.TournamentTypeAmericanoFixed:
		return &AmericanoFixedGenerator{Courts: courts, MaxPoints: maxPoints}
	case domain.TournamentTypeMixAmericano:
		return &MixAmericanoGenerator{Courts: courts, MaxPoints: maxPoints}
	case domain.TournamentTypeMixAmericanoFixed:
		return &MixAmericanoFixedGenerator{Courts: courts, MaxPoints: maxPoints}
	default:
		return nil
	}
}
