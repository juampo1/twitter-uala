package domain

import (
	"twitter-uala/internal/domain/user"
	"twitter-uala/internal/interfaces"
	"twitter-uala/repositories"
)

type Services struct {
	UserService interfaces.UserService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		UserService: user.NewUserService(repos.UserRepository),
	}
}
