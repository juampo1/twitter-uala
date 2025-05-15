package user

import (
	"context"
	"fmt"
	"twitter-uala/internal/domain/user/models"
	"twitter-uala/internal/interfaces"
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

func (s *userService) CreateTweet(ctx context.Context, content, userID string) error {
	fmt.Printf("content: %s\n", content)
	return nil
}
