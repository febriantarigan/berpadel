package service

import (
	"context"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) SearchByName(ctx context.Context, keyword string) ([]*domain.User, error) {
	return s.userRepo.SearchByName(ctx, keyword)
}

func (s *UserService) List(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.List(ctx)
}
