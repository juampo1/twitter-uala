package domain

import (
	"twitter-uala/internal/interfaces"
	"twitter-uala/repositories"
)

type Services struct {
	UserRepository interfaces.UserRepository
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		UserRepository: repos.UserRepository,
	}
}
