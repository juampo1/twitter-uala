package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:50;unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
