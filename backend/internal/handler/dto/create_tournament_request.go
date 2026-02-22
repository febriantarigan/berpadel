package dto

type CreateTournamentRequest struct {
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	Status    string          `json:"status"`
	Location  string          `json:"location"`
	Season    string          `json:"season"`
	Courts    []string        `json:"courts"`
	MaxPoints int             `json:"max_points"`
	Players   []PlayerRequest `json:"players"`
}

type PlayerRequest struct {
	UserID *string `json:"user_id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Gender *string `json:"gender,omitempty"`
}
