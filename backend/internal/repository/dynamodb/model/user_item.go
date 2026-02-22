package model

type UserItem struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`

	Name      string `dynamodbav:"name"`
	Gender    string `dynamodbav:"gender"`
	CreatedAt string `dynamodbav:"created_at"`
	UpdatedAt string `dynamodbav:"updated_at"`
}
