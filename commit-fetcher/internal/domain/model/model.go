package model

import "time"

type Commit struct {
	ID           int    `gorm:"primaryKey"`
	SHA          string `gorm:"unique;not null"`
	RepositoryID int    `gorm:"index"`
	Message      string `gorm:"not null"`
	Author       string
	Date         time.Time
	URL          string `gorm:"not null"`
}
