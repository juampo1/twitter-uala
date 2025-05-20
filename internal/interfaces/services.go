package interfaces

import (
	"context"
	followModels "twitter-uala/internal/domain/follow/models"
	tweetModels "twitter-uala/internal/domain/tweet/models"
	userModels "twitter-uala/internal/domain/user/models"
)

type (
	UserService interface {
		FindUser(ctx context.Context, userID string) (*userModels.User, error)
		CreateTweet(ctx context.Context, content, userID string) error
		FollowUser(ctx context.Context, followerID, followedUserID string) error
		GetTimeline(ctx context.Context, userID string) (*[]tweetModels.Tweet, error)
	}

	TweetService interface {
		CreateTweet(ctx context.Context, content, userID string) (*tweetModels.Tweet, error)
		GetTweetsByUserIDs(ctx context.Context, followedUsers []followModels.Follow) (*[]tweetModels.Tweet, error)
	}
)
