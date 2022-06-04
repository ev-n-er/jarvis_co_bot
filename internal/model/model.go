package model

import "time"

type Visit struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement:false"`
	Date      time.Time `gorm:"primaryKey;"`
	Type      int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
