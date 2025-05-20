package domain

import (
	"twitter-uala/internal/domain/tweet"
	"twitter-uala/internal/domain/user"
	"twitter-uala/internal/interfaces"
	"twitter-uala/repositories"
)

type Services struct {
	UserService  interfaces.UserService
	TweetService interfaces.TweetService
}

func NewServices(repos *repositories.Repositories) *Services {
	tweetService := tweet.NewTweetService(repos.TweetRepository)
	return &Services{
		UserService:  user.NewUserService(repos.UserRepository, tweetService),
		TweetService: tweetService,
	}
}
