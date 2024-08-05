package repository

import (
	"commit-monitor/internal/model"
	"time"

	"gorm.io/gorm"
)

type commitRepository struct {
	db *gorm.DB
}

func NewCommitRepository(db *gorm.DB) CommitRepository {
	return &commitRepository{db: db}
}

func (r *commitRepository) Save(commit *model.Commit) error {
	return r.db.Create(commit).Error
}

func (r *commitRepository) GetBySHA(sha string) (*model.Commit, error) {
	var commit model.Commit
	err := r.db.Where("sha = ?", sha).First(&commit).Error
	if err != nil {
		return nil, err
	}
	return &commit, nil
}

func (r *commitRepository) GetByRepositoryIDAndDateRange(repoID int, since, until time.Time) ([]model.Commit, error) {
	var commits []model.Commit
	err := r.db.Where("repository_id = ? AND date >= ? AND date <= ?", repoID, since, until).Find(&commits).Error
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (r *commitRepository) DeleteByRepositoryIDAndDate(repoID int, since time.Time) error {
	return r.db.Where("repository_id = ? AND date >= ?", repoID, since).Delete(&model.Commit{}).Error
}
