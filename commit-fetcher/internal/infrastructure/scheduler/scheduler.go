package scheduler

import (
	"commit_fetcher/internal/service"
	"log"
	"time"
)

func StartFetcherService(fs *service.FetchService) {
	log.Printf("commit fecher cronjob about to start")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := fs.FetchCommits()
			if err != nil {
				log.Printf("Error fetching commits: %v", err)
			}
			log.Printf("commit fecher cronjob started")
		}
	}
}
