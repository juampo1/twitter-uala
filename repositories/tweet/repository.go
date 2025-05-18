package user

import (
	"context"
	"fmt"
	followModels "twitter-uala/internal/domain/follow/models"
	"twitter-uala/internal/domain/tweet/models"
	"twitter-uala/internal/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.TweetRepository {
	return &repository{db: db}
}

func (r *repository) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	if err := r.db.Create(tweet).Error; err != nil {
		return nil, fmt.Errorf("error creating tweet: %w", err)
	}

	return tweet, nil
}

func (r *repository) GetTweetsByUserIDs(ctx context.Context, followedUsers []followModels.Follow) (*[]models.Tweet, error) {
	var tweets []models.Tweet
	var followedUserIDs []string

	for _, followedUser := range followedUsers {
		followedUserIDs = append(followedUserIDs, followedUser.FollowedID)
	}

	err := r.db.WithContext(ctx).Where("user_id IN (?)", followedUserIDs).Find(&tweets).Error
	if err != nil {
		return nil, fmt.Errorf("error getting tweets: %w", err)
	}

	return &tweets, nil
}
