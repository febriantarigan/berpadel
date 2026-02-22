package service

import (
	"context"
	"fmt"
	"time"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository"
	"github.com/oklog/ulid/v2"
)

type CreateTournamentInput struct {
	Name      string
	Type      domain.TournamentType
	Status    domain.TournamentStatus
	Location  string
	Season    string
	MaxPoints int
	Players   []PlayerInput
	Courts    []string
}

type PlayerInput struct {
	UserID *string
	Name   *string
	Gender domain.Gender
}

type TournamentService struct {
	userRepo        repository.UserRepository
	tournamentRepo  repository.TournamentRepository
	matchRepo       repository.MatchRepository
	leaderboardRepo repository.LeaderboardRepository
}

func NewTournamentService(userRepo repository.UserRepository, tournamentRepo repository.TournamentRepository, matchRepo repository.MatchRepository, leaderboardRepo repository.LeaderboardRepository) *TournamentService {
	return &TournamentService{
		userRepo:        userRepo,
		tournamentRepo:  tournamentRepo,
		matchRepo:       matchRepo,
		leaderboardRepo: leaderboardRepo,
	}
}

type CreateTournamentResult struct {
	Tournament *domain.Tournament
	Rounds     []*RoundSchedule
}

func (s *TournamentService) CreateTournament(
	ctx context.Context,
	input CreateTournamentInput,
) (*CreateTournamentResult, error) {
	playerIDs := make([]string, 0, len(input.Players))

	var existingIDs []string
	var newPlayers []*domain.User

	for _, p := range input.Players {
		if p.UserID != nil {
			existingIDs = append(existingIDs, *p.UserID)
			playerIDs = append(playerIDs, *p.UserID)
		} else if p.Name != nil {
			newPlayer := &domain.User{
				ID:        ulid.Make().String(),
				Name:      *p.Name,
				Gender:    p.Gender,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			newPlayers = append(newPlayers, newPlayer)
			playerIDs = append(playerIDs, newPlayer.ID)
		}
	}

	fetchedUsers, err := s.userRepo.GetByIDs(ctx, existingIDs)
	if err != nil {
		return nil, err
	}

	// Build id->user map and preserve playerIDs order for generators
	idToUser := make(map[string]*domain.User)
	for _, u := range fetchedUsers {
		idToUser[u.ID] = u
	}
	for _, u := range newPlayers {
		idToUser[u.ID] = u
	}
	users := make([]*domain.User, 0, len(playerIDs))
	for _, id := range playerIDs {
		if u, ok := idToUser[id]; ok {
			users = append(users, u)
		}
	}

	// Validate before any database writes
	if len(input.Courts) == 0 {
		return nil, fmt.Errorf("at least one court is required")
	}
	if len(users) < 4 {
		return nil, fmt.Errorf("at least 4 players are required")
	}
	gen := GetGenerator(input.Type, input.Courts, input.MaxPoints)
	if gen == nil {
		return nil, fmt.Errorf("unsupported tournament type: %s", input.Type)
	}

	//prepare tournament
	tournament := &domain.Tournament{
		ID:        ulid.Make().String(),
		Name:      input.Name,
		Type:      input.Type,
		Status:    input.Status,
		Season:    input.Season,
		PlayerIDs: playerIDs,
		MaxPoints: input.MaxPoints,
		Courts:    input.Courts,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(newPlayers) > 0 {
		if err := s.tournamentRepo.CreateWithNewUsers(ctx, tournament, newPlayers); err != nil {
			return nil, err
		}
	} else {
		if err := s.tournamentRepo.Create(ctx, tournament); err != nil {
			return nil, err
		}
	}

	//generate matches
	schedules, err := gen.Generate(tournament, users)
	if err != nil {
		return nil, err
	}

	matches := []*domain.Match{}
	for _, s := range schedules {
		matches = append(matches, s.Matches...)
	}

	if err := s.matchRepo.BatchCreate(ctx, matches); err != nil {
		return nil, err
	}

	return &CreateTournamentResult{
		Tournament: tournament,
		Rounds:     schedules,
	}, nil
}
