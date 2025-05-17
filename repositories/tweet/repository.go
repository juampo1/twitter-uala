package user

import (
	"context"
	"fmt"
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
