package model

type TournamentItem struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`

	Name      string   `dynamodbav:"name"`
	Type      string   `dynamodbav:"type"`
	Status    string   `dynamodbav:"status"`
	Location  string   `dynamodbav:"location"`
	Season    string   `dynamodbav:"season"`
	MaxPoints int      `dynamodbav:"max_points"`
	Courts    []string `dynamodbav:"courts"`
	PlayerIDs []string `dynamodbav:"player_ids"`
	CreatedAt string   `dynamodbav:"created_at"`
	UpdatedAt string   `dynamodbav:"updated_at"`

	GSI1PK string `dynamodbav:"GSI1PK"`
	GSI1SK string `dynamodbav:"GSI1SK"`
}
