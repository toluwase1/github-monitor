package model

import "time"

type Commit struct {
	ID             int    `gorm:"primaryKey"`
	RepositoryID   int    `gorm:"index"`
	Message        string `gorm:"not null"`
	Author         string
	Date           time.Time
	URL            string `gorm:"unique;not null"`
	SHA            string `gorm:"unique;not null"`
	RepositoryName string
}
