package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	GithubToken string
	StartDate   time.Time
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	startDate, err := time.Parse("2006-01-02", os.Getenv("START_DATE"))
	if err != nil {
		log.Fatalf("Invalid START_DATE format. Use YYYY-MM-DD.")
	}

	return Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		GithubToken: os.Getenv("GITHUB_TOKEN"),
		StartDate:   startDate,
	}
}
