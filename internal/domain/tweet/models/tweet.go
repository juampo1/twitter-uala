package models

import (
	"time"
)

type Tweet struct {
	ID        uint      `gorm:"primaryKey;unique;autoIncrement"`
	UserID    string    `gorm:"index;not null"`
	Content   string    `gorm:"size:280;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
