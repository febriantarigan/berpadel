package domain

import (
	"fmt"
	"time"
)

type TournamentType string
type TournamentStatus string

const (
	TournamentTypeAmericano         TournamentType = "americano"          // Rotating partners, any gender
	TournamentTypeAmericanoFixed    TournamentType = "americano_fixed"    // Fixed teams, any gender
	TournamentTypeMixAmericano      TournamentType = "mix_americano"      // Rotating partners, mixed M-F
	TournamentTypeMixAmericanoFixed TournamentType = "mix_americano_fixed" // Fixed teams, mixed M-F
	TournamentTypeMexicano          TournamentType = "mexicano"
)

const (
	TournamentStatusDraft     TournamentStatus = "draft"
	TournamentStatusOngoing   TournamentStatus = "ongoing"
	TournamentStatusCompleted TournamentStatus = "completed"
)

type Tournament struct {
	ID        string
	Name      string
	Location  string
	Type      TournamentType
	Status    TournamentStatus
	Season    string
	PlayerIDs []string
	Courts    []string
	MaxPoints int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ParseTournamentType(t string) (TournamentType, error) {
	switch t {
	case "americano":
		return TournamentTypeAmericano, nil
	case "americano_fixed":
		return TournamentTypeAmericanoFixed, nil
	case "mix_americano":
		return TournamentTypeMixAmericano, nil
	case "mix_americano_fixed":
		return TournamentTypeMixAmericanoFixed, nil
	case "mexicano":
		return TournamentTypeMexicano, nil
	default:
		return "", fmt.Errorf("invalid tournament type: %s", t)
	}
}

func ParseTournamentStatus(t string) (TournamentStatus, error) {
	switch t {
	case "draft":
		return TournamentStatusDraft, nil
	case "ongoing":
		return TournamentStatusOngoing, nil
	case "completed":
		return TournamentStatusCompleted, nil
	default:
		return "", fmt.Errorf("invalid tournament status: %s", t)
	}
}
