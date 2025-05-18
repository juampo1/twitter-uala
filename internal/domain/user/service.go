package user

import (
	"context"
	"errors"
	"fmt"
	"twitter-uala/internal/domain/user/models"
	"twitter-uala/internal/interfaces"
)

type userService struct {
	repo         interfaces.UserRepository
	tweetService interfaces.TweetService
}

func NewUserService(repo interfaces.UserRepository, tweetService interfaces.TweetService) interfaces.UserService {
	return &userService{
		repo:         repo,
		tweetService: tweetService,
	}
}

func (s *userService) FindUser(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.FindUserByID(ctx, userID)
}

func (s *userService) CreateTweet(ctx context.Context, content, userID string) error {
	fmt.Printf("content: %s\n", content)

	//check if user exists
	user, err := s.FindUser(ctx, userID)
	if err != nil {
		return err
	}

	_, err = s.tweetService.CreateTweet(ctx, content, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) FollowUser(ctx context.Context, followerID, followedUserID string) error {
	if followerID == followedUserID {
		return errors.New("a user cannot follow himself")
	}

	_, err := s.FindUser(ctx, followerID)
	if err != nil {
		return err
	}

	_, err = s.FindUser(ctx, followedUserID)
	if err != nil {
		return err
	}

	//TODO: check if user is already followed by the follower

	err = s.repo.FollowUser(ctx, followerID, followedUserID)
	if err != nil {
		return err
	}

	return nil
}
