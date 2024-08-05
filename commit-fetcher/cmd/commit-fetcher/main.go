package main

import (
	"commit_fetcher/internal/domain/repository"
	"commit_fetcher/internal/infrastructure/config"
	"commit_fetcher/internal/interface/github"
	"commit_fetcher/internal/service"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize GORM database connection
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize GitHub client
	githubClient := github.NewClient(cfg.GithubToken)

	// Initialize repositories
	commitRepo := repository.NewCommitRepository(db)

	// Initialize fetch service
	fetchService := service.NewFetchService(commitRepo, githubClient, cfg)

	// Fetch commits
	err = fetchService.FetchCommits()
	if err != nil {
		log.Fatalf("Error fetching commits: %v", err)
	}

	//// Send signal to RabbitMQ after fetching
	//publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQURL, cfg.RepoName)
	//if err != nil {
	//	log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	//}
	//defer publisher.Close()
	//
	//err = publisher.PublishSignal("FetchComplete")
	//if err != nil {
	//	log.Fatalf("Failed to send signal: %v", err)
	//}
}
