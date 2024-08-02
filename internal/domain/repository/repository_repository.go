package repository

import (
	"github-monitor/internal/domain/model"
	"gorm.io/gorm"
)

type RepositoryRepository interface {
	Save(repo *model.Repository) error
	GetByName(name string) (*model.Repository, error)
}

type repositoryRepository struct {
	db *gorm.DB
}

func NewRepositoryRepository(db *gorm.DB) RepositoryRepository {
	return &repositoryRepository{db: db}
}

func (r *repositoryRepository) Save(repo *model.Repository) error {
	return r.db.Create(repo).Error
}

func (r *repositoryRepository) GetByName(name string) (*model.Repository, error) {
	var repo model.Repository
	if err := r.db.Where("name = ?", name).First(&repo).Error; err != nil {
		return nil, err
	}
	return &repo, nil
}
