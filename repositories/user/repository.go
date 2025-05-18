package user

import (
	"context"
	"fmt"
	followModels "twitter-uala/internal/domain/follow/models"
	"twitter-uala/internal/domain/user/models"
	userModels "twitter-uala/internal/domain/user/models"
	"twitter-uala/internal/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.UserRepository {
	return &repository{db: db}
}

func (r *repository) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user userModels.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("could not find user: %w", err)
	}
	return &user, nil
}

func (r *repository) FollowUser(ctx context.Context, followerID, followedUserID string) error {
	follower := &followModels.Follow{
		UserID:     followerID,
		FollowedID: followedUserID,
	}

	if err := r.db.WithContext(ctx).Create(follower).Error; err != nil {
		return fmt.Errorf("could not follow user: %w", err)
	}

	return nil
}

func (r *repository) GetFollowedUsers(ctx context.Context, userID string) ([]followModels.Follow, error) {
	var followedUsers []followModels.Follow
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&followedUsers).Error; err != nil {
		return nil, fmt.Errorf("could not get followed users: %w", err)
	}
	return followedUsers, nil
}
