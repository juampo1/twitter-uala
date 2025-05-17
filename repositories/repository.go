package repositories

import (
	"twitter-uala/internal/interfaces"
	tweetRepository "twitter-uala/repositories/tweet"
	userRepository "twitter-uala/repositories/user"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository  interfaces.UserRepository
	TweetRepository interfaces.TweetRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:  userRepository.NewRepository(db),
		TweetRepository: tweetRepository.NewRepository(db),
	}
}
