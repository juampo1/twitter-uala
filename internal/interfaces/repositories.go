package interfaces

import (
	"context"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"
)

type (
	UserRepository interface {
		FindUserByID(ctx context.Context, userID string) (*userModels.User, error)
	}
	TweetRepository interface {
		CreateTweet(ctx context.Context, tweet *tweetModels.Tweet) (*tweetModels.Tweet, error)
	}
)
