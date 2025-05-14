package models

type Follow struct {
	FollowerID uint `gorm:"not null"`
	FollowedID uint `gorm:"not null"`
	Follower   User `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed   User `gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}
