package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/febriantarigan/berpadel/internal/config"
)

type BaseRepository struct {
	DB  *dynamodb.Client
	cfg *config.DynamoDB
}

func NewBaseRepository(db *dynamodb.Client) *BaseRepository {
	return &BaseRepository{
		DB:  db,
		cfg: config.GetDynamoDB(),
	}
}
