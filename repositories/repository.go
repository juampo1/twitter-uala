package repositories

import (
	"twitter-uala/internal/domain/interfaces"
	"twitter-uala/repositories/user"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository interfaces.UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: user.NewRepository(db),
	}
}
