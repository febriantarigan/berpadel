package model

type LeaderboardItem struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`

	UserID  string `dynamodbav:"UserID"`
	Matches int    `dynamodbav:"Matches"`
	Wins    int    `dynamodbav:"Wins"`
	Draws   int    `dynamodbav:"Draws"`
	Losses  int    `dynamodbav:"Losses"`
	Points  int    `dynamodbav:"Points"`
}
