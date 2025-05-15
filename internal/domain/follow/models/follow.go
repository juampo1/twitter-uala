package models

import "twitter-uala/internal/domain/user/models"

type Follow struct {
	FollowerID uint        `gorm:"not null"`
	FollowedID uint        `gorm:"not null"`
	Follower   models.User `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed   models.User `gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}
