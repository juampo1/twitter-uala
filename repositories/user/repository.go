package user

import (
	"context"
	"fmt"
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
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("error al encontrar el usuario: %w", err)
	}
	return &user, nil
}
