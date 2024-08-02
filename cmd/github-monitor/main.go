package main

import (
	"github-monitor/internal/domain/model"
	"github-monitor/internal/domain/repository"
	"github-monitor/internal/infrastructure/config"
	"github-monitor/internal/infrastructure/router"
	"github-monitor/internal/infrastructure/scheduler"
	"github-monitor/internal/interface/controller"
	"github-monitor/internal/interface/github"
	"github-monitor/internal/usecase"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var repoName = "chromium/chromium"

func main() {
	cfg := config.LoadConfig()

	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Repository{}, &model.Commit{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	githubClient := github.NewGithubClient(cfg.GithubToken)

	repoRepo := repository.NewRepositoryRepository(db)
	commitRepo := repository.NewCommitRepository(db)

	repoUsecase := usecase.NewRepositoryUsecase(repoRepo, githubClient)
	commitUsecase := usecase.NewCommitUsecase(commitRepo, repoRepo, githubClient)

	ctrl := controller.NewController(commitUsecase, repoUsecase)

	getRepoDetails(err, repoRepo, githubClient)

	scheduler.StartCommitScheduler(repoName, commitUsecase)

	r := router.NewRouter(ctrl)
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getRepoDetails(err error, repoRepo repository.RepositoryRepository, githubClient *github.GithubClient) {
	repoDetails, err := repoRepo.GetByName(repoName)
	if err != nil || repoDetails == nil {
		repoDetails, err = githubClient.GetRepository(repoName)
		if err != nil {
			log.Fatalf("Failed to fetch repository details: %v", err)
		}
		if err := repoRepo.Save(repoDetails); err != nil {
			log.Fatalf("Failed to save repository details: %v", err)
		}
	} else {
		log.Println("Repository details already exist in the database.")
	}
}
