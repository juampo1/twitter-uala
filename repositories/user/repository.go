package user

import (
	"context"
	"twitter-uala/internal/domain/user/models"
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
	return nil, nil
}
