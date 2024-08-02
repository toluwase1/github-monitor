package model

import "time"

type Repository struct {
	ID              int    `gorm:"primaryKey"`
	Name            string `gorm:"unique;not null"`
	Description     string
	URL             string
	Language        string
	ForksCount      int
	StarsCount      int
	OpenIssuesCount int
	WatchersCount   int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
