package main

import (
	"commit-monitor/internal/infrastructure/config"
	rabbitmq "commit-monitor/internal/infrastructure/rabbit"
	"commit-monitor/internal/infrastructure/scheduler"
	"commit-monitor/internal/interface/repository"
	"commit-monitor/internal/service"
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

	// Initialize RabbitMQ consumer
	consumer, err := rabbitmq.NewConsumer(cfg.RabbitMQURL, cfg.RepoName)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ consumer: %v", err)
	}
	defer consumer.Close()

	// Wait for signal to start monitoring
	msgs, err := consumer.ConsumeSignals()
	if err != nil {
		log.Fatalf("Failed to consume signals: %v", err)
	}

	for msg := range msgs {
		if string(msg.Body) == "FetchComplete" {
			// Start monitoring new commits
			monitorService := scheduler.NewMonitorService(commitRepo, githubClient)
			go func() {
				err = monitorService.MonitorNewCommits()
				if err != nil {
					log.Printf("an error occured while monitoring commits %v", err)
				}
			}()
			break
		}
	}

	select {} // Block main goroutine
}
