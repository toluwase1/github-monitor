package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	GithubToken string
	SinceDate   time.Time
	UntilDate   time.Time
	RabbitMQURL string
	QueueName   string
	RepoName    string
}

func LoadConfig() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	sinceDate, err := time.Parse("2006-01-02", os.Getenv("SINCE_DATE"))
	if err != nil {
		panic("Invalid SINCE_DATE format. Use YYYY-MM-DD.")
	}
	untilDate, err := time.Parse("2006-01-02", os.Getenv("UNTIL_DATE"))
	if err != nil {
		panic("Invalid UNTIL_DATE format. Use YYYY-MM-DD.")
	}

	return Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		GithubToken: os.Getenv("GITHUB_TOKEN"),
		SinceDate:   sinceDate,
		UntilDate:   untilDate,
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
		QueueName:   os.Getenv("QUEUE_NAME"),
		RepoName:    os.Getenv("REPO_NAME"),
	}
}
