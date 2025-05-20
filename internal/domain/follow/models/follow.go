package models

type Follow struct {
	ID         string `gorm:"primaryKey;unique;autoIncrement"`
	UserID     string `gorm:"index;not null"`
	FollowedID string `gorm:"index;not null"`
}
