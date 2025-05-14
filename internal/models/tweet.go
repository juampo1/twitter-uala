package models

import "time"

type Tweet struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"size:280;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
