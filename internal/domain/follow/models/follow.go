package models

type Follow struct {
	ID         string `gorm:"primaryKey;autoIncrement"`
	UserID     string `gorm:"index;not null"`
	FollowedID string `gorm:"index;not null"`
}
