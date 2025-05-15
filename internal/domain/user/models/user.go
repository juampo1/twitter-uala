package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
