package repositories

import (
	"twitter-uala/internal/interfaces"
	tweetRepository "twitter-uala/repositories/tweet"
	userRepository "twitter-uala/repositories/user"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository  interfaces.UserRepository
	TweetRepository interfaces.TweetRepository
}

func NewRepositories(db *gorm.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		UserRepository:  userRepository.NewRepository(db, redis),
		TweetRepository: tweetRepository.NewRepository(db, redis),
	}
}
