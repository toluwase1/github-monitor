package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	GithubToken string
	RabbitMQURL string
	QueueName   string
	RepoName    string
}

func LoadConfig() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		GithubToken: os.Getenv("GITHUB_TOKEN"),
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
		QueueName:   os.Getenv("QUEUE_NAME"),
		RepoName:    os.Getenv("REPO_NAME"),
	}
}
