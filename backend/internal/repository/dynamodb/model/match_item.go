package model

type MatchItem struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`

	MatchID          string   `dynamodbav:"match_id"`
	Round            int      `dynamodbav:"round"`
	Court            string   `dynamodbav:"court"`
	TeamA            []string `dynamodbav:"team_a"`
	TeamB            []string `dynamodbav:"team_b"`
	WaitingPlayerIDs []string `dynamodbav:"waiting_player_ids"`
	TeamAScore       int      `dynamodbav:"team_a_score"`
	TeamBScore       int      `dynamodbav:"team_b_score"`
	Status           string   `dynamodbav:"status"`
	CreatedAt        string   `dynamodbav:"created_at"`
	UpdatedAt        string   `dynamodbav:"updated_at"`
}
