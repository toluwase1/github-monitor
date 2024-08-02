// internal/infrastructure/scheduler/scheduler.go

package scheduler

import (
	"github-monitor/internal/usecase"
	"log"
	"time"
)

func StartCommitScheduler(repoName string, commitUsecase usecase.CommitUsecase) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Checking for new commits...")
				err := checkForNewCommits(repoName, commitUsecase)
				if err != nil {
					log.Printf("Error checking for new commits: %v", err)
				}
			}
		}
	}()
}

func checkForNewCommits(repoName string, commitUsecase usecase.CommitUsecase) error {
	// Fetch the latest commits from the GitHub API
	commits, err := commitUsecase.GetCommitsByRepositoryName(repoName)
	if err != nil {
		return err
	}

	// Process commits (e.g., save to DB if not already present)
	for _, commit := range commits {
		err := commitUsecase.SaveCommitIfNotExists(&commit)
		if err != nil {
			log.Printf("Error saving commit: %v", err)
		}
	}

	return nil
}
