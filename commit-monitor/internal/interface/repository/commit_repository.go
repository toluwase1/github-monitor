package repository

import (
	"commit-monitor/internal/model"
	"time"
)

type CommitRepository interface {
	Save(commit *model.Commit) error
	GetBySHA(sha string) (*model.Commit, error)
	GetByRepositoryIDAndDateRange(repoID int, since, until time.Time) ([]model.Commit, error)
	DeleteByRepositoryIDAndDate(repoID int, since time.Time) error
}
