package service

import (
	"commit_fetcher/internal/domain/repository"
	"commit_fetcher/internal/infrastructure/config"
	"commit_fetcher/internal/infrastructure/rabbit"
	"commit_fetcher/internal/interface/github"
	"log"
)

type FetchService struct {
	commitRepo   repository.CommitRepository
	githubClient *github.Client
	config       config.Config
}

func NewFetchService(commitRepo repository.CommitRepository, githubClient *github.Client, config config.Config) *FetchService {
	return &FetchService{commitRepo, githubClient, config}
}

func (fs *FetchService) FetchCommits() error {
	// Fetch commits from GitHub API
	commits, err := fs.githubClient.GetCommits(fs.config.SinceDate, fs.config.UntilDate)
	if err != nil {
		return err
	}

	for _, commit := range commits {
		existingCommit, _ := fs.commitRepo.GetBySHA(commit.SHA)
		if existingCommit == nil {
			if err := fs.commitRepo.Save(&commit); err != nil {
				log.Printf("Failed to save commit: %v", err)
			}
		}
	}

	// Send signal to RabbitMQ
	publisher, err := rabbitmq.NewPublisher(fs.config.RabbitMQURL, fs.config.QueueName)
	if err != nil {
		return err
	}
	defer publisher.Close()

	err = publisher.PublishSignal("FetchComplete")
	if err != nil {
		return err
	}

	return nil
}
