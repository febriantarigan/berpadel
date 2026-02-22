package domain

type Leaderboard struct {
	UserID  string
	Matches int
	Wins    int
	Draws   int
	Losses  int
	Points  int
}
