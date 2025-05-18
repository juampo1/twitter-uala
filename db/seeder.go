package db

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	follow "twitter-uala/internal/domain/follow/models"
	tweet "twitter-uala/internal/domain/tweet/models"
	"twitter-uala/internal/domain/user/models"

	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) Seed() {
	s.DeleteAll()
	users := []models.User{
		{
			ID:       "1",
			Username: "@juanito",
		},
		{
			ID:       "2",
			Username: "@mary",
		},
		{
			ID:       "3",
			Username: "@carlitos",
		},
		{
			ID:       "4",
			Username: "@anita",
		},
		{
			ID:       "5",
			Username: "@pedrito",
		},
		{
			ID:       "6",
			Username: "@sofi",
		},
		{
			ID:       "7",
			Username: "@luisito",
		},
		{
			ID:       "8",
			Username: "@lau",
		},
		{
			ID:       "10",
			Username: "@mike",
		},
		{
			ID:       "11",
			Username: "@carmi",
		},
		{
			ID:       "12",
			Username: "@joseto",
		},
		{
			ID:       "13",
			Username: "@isa",
		},
		{
			ID:       "14",
			Username: "@dieguito",
		},
		{
			ID:       "15",
			Username: "@vale",
		},
		{
			ID:       "16",
			Username: "@sebas",
		},
		{
			ID:       "17",
			Username: "@gabi",
		},
		{
			ID:       "18",
			Username: "@manu",
		},
		{
			ID:       "19",
			Username: "@cami",
		},
		{
			ID:       "20",
			Username: "@rodrigo",
		},
	}

	// Create a map to store follow relationships and avoid duplicates
	followMap := make(map[string]map[string]bool)

	// Initialize the map
	for _, user := range users {
		followMap[user.ID] = make(map[string]bool)
	}

	// Assign followers to each user
	for _, user := range users {
		// Generate a random number of followers between 5 and 10
		numFollowers := rand.Intn(6) + 5

		// insert user into the database
		if err := s.db.Create(&user).Error; err != nil {
			log.Fatalf("Error inserting user %s: %v", user.Username, err)
		}

		// Get a list of other users (excluding the current user)
		var otherUsers []models.User
		for _, u := range users {
			if u.ID != user.ID {
				otherUsers = append(otherUsers, u)
			}
		}

		// Shuffle the list of other users
		rand.Shuffle(len(otherUsers), func(i, j int) {
			otherUsers[i], otherUsers[j] = otherUsers[j], otherUsers[i]
		})

		// Adjust if the number of followers exceeds the number of available users
		if numFollowers > len(otherUsers) {
			numFollowers = len(otherUsers)
		}

		// Select the first 'numFollowers' users as followers
		followers := otherUsers[:numFollowers]

		// Assign followers to the current user
		for _, follower := range followers {
			// Check if the relationship already exists to avoid duplicates
			if !followMap[user.ID][follower.ID] {
				// Create the follow relationship in the database
				follow := follow.Follow{
					UserID:     user.ID,
					FollowedID: follower.ID,
				}
				if err := s.db.Create(&follow).Error; err != nil {
					log.Printf("Error al insertar seguimiento de %s a %s: %v", follower.Username, user.Username, err)
				}

				// Mark the relationship as existing
				followMap[user.ID][follower.ID] = true
			}
		}
	}

	// Create tweets for each user
	for _, user := range users {
		// Generate a random number of tweets between 1 and 5
		numTweets := rand.Intn(5) + 1

		// Create tweets for the current user
		for i := 0; i < numTweets; i++ {
			tweet := tweet.Tweet{
				UserID:    user.ID,
				Content:   fmt.Sprintf("Tweet %d from %s", i+1, user.Username),
				CreatedAt: randomTimeFromNowToOneMonthAgo(),
			}
			if err := s.db.Create(&tweet).Error; err != nil {
				log.Printf("Error inserting tweet for user %s: %v", user.Username, err)
			}
		}

	}
}

func randomTimeFromNowToOneMonthAgo() time.Time {
	now := time.Now()
	// Define the duration of one month (approximately 30 days)
	oneMonthAgo := now.AddDate(0, -1, 0)
	// Calculate the difference in seconds
	diff := now.Unix() - oneMonthAgo.Unix()
	// Generate a random number of seconds within that range
	randomSeconds := rand.Int63n(diff)
	// Subtract the random seconds from now
	randomTime := now.Add(-time.Duration(randomSeconds) * time.Second)
	return randomTime
}

func (s *Seeder) DeleteAll() {
	err := s.db.Exec("DELETE FROM users").Error
	if err != nil {
		log.Fatalf("Error deleting users table")
	}

	err = s.db.Exec("DELETE FROM follows").Error
	if err != nil {
		log.Fatalf("Error deleting follows table")
	}

	err = s.db.Exec("DELETE FROM tweets").Error
	if err != nil {
		log.Fatalf("Error deleting tweets table")
	}
}
