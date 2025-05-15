package domain

import (
	"twitter-uala/internal/domain/interfaces"
	"twitter-uala/internal/repositories"
)

type Services struct {
	UserRepository interfaces.UserRepository
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		UserRepository: repos.UserRepository,
	}
}
