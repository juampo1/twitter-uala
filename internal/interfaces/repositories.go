package interfaces

import (
	"context"
	followModels "twitter-uala/internal/domain/follow/models"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"
)

type (
	UserRepository interface {
		FindUserByID(ctx context.Context, userID string) (*userModels.User, error)
		FollowUser(ctx context.Context, followerID, followedUserID string) error
		GetFollowedUsers(ctx context.Context, userID string) ([]followModels.Follow, error)
	}
	TweetRepository interface {
		CreateTweet(ctx context.Context, tweet *tweetModels.Tweet) (*tweetModels.Tweet, error)
		GetTweetsByUserIDs(ctx context.Context, followedUsers []followModels.Follow) (*[]tweetModels.Tweet, error)
	}
)
