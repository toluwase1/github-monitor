package usecase

import (
	"github-monitor/internal/domain/model"
	"github-monitor/internal/domain/repository"
	"github-monitor/internal/interface/github"
	"time"
)

type CommitUsecase interface {
	GetCommitsByRepositoryName(name string) ([]model.Commit, error)
	SaveCommitIfNotExists(commit *model.Commit) error
	ResetCollection(name string, startDate time.Time) error
}

type commitUsecase struct {
	commitRepo     repository.CommitRepository
	repositoryRepo repository.RepositoryRepository
	githubClient   *github.GithubClient
}

func NewCommitUsecase(commitRepo repository.CommitRepository, repoRepo repository.RepositoryRepository, githubClient *github.GithubClient) CommitUsecase {
	return &commitUsecase{
		commitRepo:     commitRepo,
		repositoryRepo: repoRepo,
		githubClient:   githubClient,
	}
}

func (uc *commitUsecase) GetCommitsByRepositoryName(name string) ([]model.Commit, error) {
	return uc.githubClient.GetCommits(name)
}

func (uc *commitUsecase) SaveCommitIfNotExists(commit *model.Commit) error {
	existingCommits, err := uc.commitRepo.GetByRepositoryID(commit.RepositoryID)
	if err != nil {
		return err
	}

	for _, existingCommit := range existingCommits {
		if existingCommit.URL == commit.URL {
			return nil
		}
	}

	return uc.commitRepo.Save(commit)
}

func (uc *commitUsecase) ResetCollection(name string, startDate time.Time) error {
	repo, err := uc.repositoryRepo.GetByName(name)
	if err != nil {
		return err
	}

	// Delete commits from database starting from the start date
	err = uc.commitRepo.DeleteByRepositoryIDAndDate(repo.ID, startDate)
	if err != nil {
		return err
	}

	// Fetch new commits from the GitHub API starting from the start date
	commits, err := uc.githubClient.GetCommits(name)
	if err != nil {
		return err
	}

	for _, commit := range commits {
		if commit.Date.After(startDate) || commit.Date.Equal(startDate) {
			err := uc.SaveCommitIfNotExists(&commit)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
