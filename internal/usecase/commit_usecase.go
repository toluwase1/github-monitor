package usecase

import (
	"github-monitor/internal/domain/model"
	"github-monitor/internal/domain/repository"
	"github-monitor/internal/interface/github"
	"log"
	"time"
)

type CommitUsecase interface {
	GetCommitsByRepositoryName(name string) ([]model.Commit, error)
	GetCommitsByRepoNameFromDB(name string) ([]model.Commit, error)
	SaveCommitIfNotExists(commit *model.Commit) error
	ResetCollection(name string, startDate time.Time) error
	GetTopAuthorsByCommitCount(n int) ([]repository.AuthorCommitCount, error)
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

func (uc *commitUsecase) GetCommitsByRepoNameFromDB(name string) ([]model.Commit, error) {
	return uc.commitRepo.GetCommitsByNameFromDB(name)
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
		log.Println("client error", err)
		return err
	}

	for _, commit := range commits {
		if commit.Date.After(startDate) || commit.Date.Equal(startDate) {
			err = uc.SaveCommitIfNotExists(&commit)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *commitUsecase) GetTopAuthorsByCommitCount(n int) ([]repository.AuthorCommitCount, error) {
	return uc.commitRepo.GetTopAuthorsByCommitCount(n)
}
