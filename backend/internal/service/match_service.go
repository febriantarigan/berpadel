package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	db "github.com/febriantarigan/berpadel/internal/repository/dynamodb"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/model"
)

type MatchService struct {
	db              *dynamodb.Client
	matchRepo       *db.MatchRepository
	leaderboardRepo *db.LeaderboardRepository
	userRepo        *db.UserRepository
}

func NewMatchService(
	db *dynamodb.Client,
	matchRepo *db.MatchRepository,
	leaderboardRepo *db.LeaderboardRepository,
	userRepo *db.UserRepository,
) *MatchService {
	return &MatchService{
		db:              db,
		matchRepo:       matchRepo,
		leaderboardRepo: leaderboardRepo,
		userRepo:        userRepo,
	}
}

func (*MatchService) GetMatch(ctx context.Context, tournamentID string, matchID string) (*model.MatchItem, error) {
	return nil, nil
}

func (*MatchService) SubmitScore(ctx context.Context, tournamentID string, matchID string) (*model.MatchItem, error) {
	return nil, nil
}
