package repository

import (
	"github-monitor/internal/domain/model"
)

type MockCommitRepository struct {
	Commits []model.Commit
	Err     error
}

func (m *MockCommitRepository) Save(commit *model.Commit) error {
	if m.Err != nil {
		return m.Err
	}
	m.Commits = append(m.Commits, *commit)
	return nil
}

func (m *MockCommitRepository) GetByRepositoryID(repositoryID int) ([]model.Commit, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	var result []model.Commit
	for _, commit := range m.Commits {
		if commit.RepositoryID == repositoryID {
			result = append(result, commit)
		}
	}
	return result, nil
}
