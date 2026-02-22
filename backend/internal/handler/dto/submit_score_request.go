package dto

type SubmitMatchScoreRequest struct {
	TeamAScore int `json:"team_a_score" binding:"required"`
	TeamBScore int `json:"team_b_score" binding:"required"`
}
