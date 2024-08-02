package repository

import (
	"github-monitor/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type CommitRepository interface {
	Save(commit *model.Commit) error
	GetByRepositoryID(repositoryID int) ([]model.Commit, error)
	DeleteByRepositoryIDAndDate(repositoryID int, startDate time.Time) error
}

type commitRepository struct {
	db *gorm.DB
}

func NewCommitRepository(db *gorm.DB) CommitRepository {
	return &commitRepository{db: db}
}
