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

	if err := commitUsecase.ResetCollection("chromium/chromium", cfg.StartDate); err != nil {
		log.Fatalf("Failed to reset collection: %v", err)
	}

	ctrl := controller.NewController(commitUsecase, repoUsecase)

	scheduler.StartCommitScheduler("chromium/chromium", commitUsecase)

	r := router.NewRouter(ctrl)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
