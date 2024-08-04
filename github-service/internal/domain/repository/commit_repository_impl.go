package repository

import (
	"github-monitor/internal/domain/model"
	"time"
)

func (r *commitRepository) Save(commit *model.Commit) error {
	return r.db.Create(commit).Error
}

func (r *commitRepository) GetByRepositoryID(repositoryID int) ([]model.Commit, error) {
	var commits []model.Commit
	if err := r.db.Where("repository_id = ?", repositoryID).Find(&commits).Error; err != nil {
		return nil, err
	}
	return commits, nil
}

func (r *commitRepository) GetCommitsByNameFromDB(name string) ([]model.Commit, error) {
	var commits []model.Commit
	if err := r.db.Where("repository_name = ?", name).Find(&commits).Error; err != nil {
		return nil, err
	}
	return commits, nil
}

func (r *commitRepository) DeleteByRepositoryIDAndDate(repositoryID int, startDate time.Time) error {
	return r.db.Where("repository_id = ? AND date >= ?", repositoryID, startDate).Delete(&model.Commit{}).Error
}

func (r *commitRepository) GetTopAuthorsByCommitCount(n int) ([]AuthorCommitCount, error) {
	var results []AuthorCommitCount
	err := r.db.Model(&model.Commit{}).
		Select("author, count(*) as count").
		Group("author").
		Order("count desc").
		Limit(n).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
