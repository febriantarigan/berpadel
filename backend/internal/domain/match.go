package domain

import "time"

type MatchStatus string
type MatchResult string

const (
	MatchScheduled MatchStatus = "scheduled"
	MatchCompleted MatchStatus = "completed"
)

const (
	MatchResultWin  MatchResult = "win"
	MatchResultLose MatchResult = "lose"
	MatchResultDraw MatchResult = "draw"
)

type Match struct {
	ID               string
	TournamentID     string
	Round            int
	Court            string
	TeamA            Team
	TeamB            Team
	WaitingPlayerIDs []string
	Status           MatchStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Team struct {
	Players []*User
	Score   int
}
