package interfaces

import (
	"context"
	"twitter-uala/internal/domain/user/models"
)

type (
	UserRepository interface {
		FindUserByID(ctx context.Context, userID string) (*models.User, error)
	}
)
