package model

import "time"

type Article struct {
	PID       int64  `gorm:"primaryKey"`
	Title     string `json:"title" form:"title" gorm:"type:varchar(255);not null"`
	Digest    string `json:"digest" form:"digest" gorm:"type:text"`
	Content   string `json:"content" form:"content" gorm:"type:text;not null"`
	Tag       string `json:"tag" form:"tag" gorm:"type:varchar(100)"`
	ReadCount int
	CreatedAt time.Time
	UpdateAt  time.Time
}
