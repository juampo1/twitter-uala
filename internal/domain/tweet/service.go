package tweet

import (
	"context"
	followModels "twitter-uala/internal/domain/follow/models"
	"twitter-uala/internal/domain/tweet/models"
	"twitter-uala/internal/interfaces"
)

type tweetService struct {
	repo interfaces.TweetRepository
}

func NewTweetService(repo interfaces.TweetRepository) interfaces.TweetService {
	return &tweetService{
		repo: repo,
	}
}

func (s *tweetService) CreateTweet(ctx context.Context, content, userID string) (*models.Tweet, error) {
	tweet := &models.Tweet{
		UserID:  userID,
		Content: content,
	}

	createdTweet, err := s.repo.CreateTweet(ctx, tweet)
	if err != nil {
		return nil, err
	}

	return createdTweet, nil
}

func (s *tweetService) GetTweetsByUserIDs(ctx context.Context, followedUsers []followModels.Follow) (*[]models.Tweet, error) {
	tweets, err := s.repo.GetTweetsByUserIDs(ctx, followedUsers)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
