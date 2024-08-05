package scheduler

import (
	"commit-monitor/internal/interface/repository"
	github "commit-monitor/internal/service"
	"log"
	"time"
)

func StartMonitorService(ms *MonitorService) {
	go func() {
		err := ms.MonitorNewCommits()
		if err != nil {
			log.Fatalf("Error monitoring new commits: %v", err)
		}
	}()
}

type MonitorService struct {
	commitRepo   repository.CommitRepository
	githubClient *github.Client
}

func NewMonitorService(commitRepo repository.CommitRepository, githubClient *github.Client) *MonitorService {
	return &MonitorService{commitRepo, githubClient}
}

func (ms *MonitorService) MonitorNewCommits() error {
	// Continuously check for new commits
	for {
		// Fetch latest commits from the GitHub API
		commits, err := ms.githubClient.GetNewCommits()
		if err != nil {
			return err
		}

		for _, commit := range commits {
			existingCommit, _ := ms.commitRepo.GetBySHA(commit.SHA)
			if existingCommit == nil {
				if err := ms.commitRepo.Save(&commit); err != nil {
					log.Printf("Failed to save commit: %v", err)
				}
			}
		}

		time.Sleep(1 * time.Hour) // Adjust the interval as needed
	}
}
