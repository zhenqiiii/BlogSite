package model

import "time"

type Article struct {
	PID       int64  `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(255);not null"`
	Digest    string `gorm:"type:text"`
	Content   string `gorm:"type:text;not null"`
	Tag       string `gorm:"type:varchar(100)"`
	ReadCount int
	CreatedAt time.Time
	UpdateAt  time.Time
}
