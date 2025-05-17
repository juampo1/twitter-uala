package interfaces

import (
	"context"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"
)

type (
	UserService interface {
		FindUser(ctx context.Context, userID string) (*userModels.User, error)
		CreateTweet(ctx context.Context, content, userID string) error
	}

	TweetService interface {
		CreateTweet(ctx context.Context, content, userID string) (*tweetModels.Tweet, error)
	}
)
