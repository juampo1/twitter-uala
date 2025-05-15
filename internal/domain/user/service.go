package user

import (
	"context"
	"twitter-uala/internal/domain/interfaces"
	"twitter-uala/internal/domain/user/models"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) FindUser(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.FindUserByID(ctx, userID)
}
