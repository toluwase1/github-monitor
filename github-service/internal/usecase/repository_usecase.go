package usecase

import (
	"github-monitor/internal/domain/model"
	"github-monitor/internal/domain/repository"
	"github-monitor/internal/interface/github"
)

type RepositoryUsecase interface {
	GetRepositoryByName(name string) (*model.Repository, error)
}

type repositoryUsecase struct {
	repositoryRepo repository.RepositoryRepository
	githubClient   *github.GithubClient
}

func NewRepositoryUsecase(repoRepo repository.RepositoryRepository, githubClient *github.GithubClient) RepositoryUsecase {
	return &repositoryUsecase{
		repositoryRepo: repoRepo,
		githubClient:   githubClient,
	}
}

func (uc *repositoryUsecase) GetRepositoryByName(name string) (*model.Repository, error) {
	repo, err := uc.repositoryRepo.GetByName(name)
	if err != nil || repo == nil {
		repo, err = uc.githubClient.GetRepository(name)
		if err != nil {
			return nil, err
		}
		err = uc.repositoryRepo.Save(repo)
		if err != nil {
			return nil, err
		}
	}
	return repo, nil
}
