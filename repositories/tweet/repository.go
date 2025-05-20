package user

import (
	"context"
	"encoding/json"
	"fmt"
	"twitter-uala/db"
	followModels "twitter-uala/internal/domain/follow/models"
	"twitter-uala/internal/domain/tweet/models"
	"twitter-uala/internal/interfaces"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) interfaces.TweetRepository {
	return &repository{db: db, redis: redis}
}

func (r *repository) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	//save it in the database
	if err := r.db.Create(tweet).Error; err != nil {
		return nil, fmt.Errorf("error creating tweet: %w", err)
	}
	//Save it in redis
	tweetData, err := db.GetTweetDataForRedis(*tweet)
	if err != nil {
		return nil, fmt.Errorf("error getting tweet data for redis: %w", err)
	}
	_, err = r.redis.LPush(ctx, fmt.Sprintf("tweet:%s", tweet.UserID), tweetData).Result()
	if err != nil {
		return nil, fmt.Errorf("error inserting tweet into redis: %w", err)
	}

	return tweet, nil
}

func (r *repository) GetTweetsByUserIDs(ctx context.Context, followedUsers []followModels.Follow) (*[]models.Tweet, error) {
	var tweetsData []models.Tweet

	for _, followedUser := range followedUsers {
		tweets, _ := r.redis.LRange(ctx, fmt.Sprintf("tweet:%s", followedUser.FollowedID), 0, -1).Result()
		for _, tweet := range tweets {
			var tw models.Tweet
			if err := json.Unmarshal([]byte(tweet), &tw); err != nil {
				return nil, fmt.Errorf("error unmarshalling tweets: %w", err)
			}
			tweetsData = append(tweetsData, tw)
		}
	}

	//Code to get followed tweets from the database
	//err := r.db.WithContext(ctx).Where("user_id IN (?)", followedUserIDs).Find(&tweets).Error
	//if err != nil {
	//	return nil, fmt.Errorf("error getting tweets: %w", err)
	//}

	return &tweetsData, nil
}
