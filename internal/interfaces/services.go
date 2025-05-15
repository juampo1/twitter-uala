package interfaces

import (
	"context"
	"twitter-uala/internal/domain/user/models"
)

type (
	UserService interface {
		FindUser(ctx context.Context, userID string) (*models.User, error)
	}
)
