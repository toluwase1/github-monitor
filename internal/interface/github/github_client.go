package github

import (
	"encoding/json"
	"fmt"
	"github-monitor/internal/domain/model"
	"io/ioutil"
	"net/http"
	"time"
)

type GithubClient struct {
	Token string
}

func NewGithubClient(token string) *GithubClient {
	return &GithubClient{Token: token}
}

// check the db for the time of the latest commit
// before fetching from that time
func (c *GithubClient) GetCommits(repoName string) ([]model.Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits", repoName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.Token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch commits, status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiCommits []struct {
		SHA    string `json:"sha"`
		Commit struct {
			Message string `json:"message"`
			Author  struct {
				Name  string    `json:"name"`
				Email string    `json:"email"`
				Date  time.Time `json:"date"`
			} `json:"author"`
		} `json:"commit"`
		URL string `json:"html_url"`
	}
	if err := json.Unmarshal(body, &apiCommits); err != nil {
		return nil, err
	}

	commits := make([]model.Commit, len(apiCommits))
	for i, apiCommit := range apiCommits {
		commits[i] = model.Commit{
			Message:        apiCommit.Commit.Message,
			Author:         apiCommit.Commit.Author.Name,
			Date:           apiCommit.Commit.Author.Date,
			URL:            apiCommit.URL,
			Sha:            apiCommit.SHA,
			RepositoryName: "chromium",
		}
	}

	return commits, nil
}

func (c *GithubClient) GetRepository(repoName string) (*model.Repository, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", repoName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.Token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repository, status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiRepo struct {
		Name            string    `json:"name"`
		Description     string    `json:"description"`
		HTMLURL         string    `json:"html_url"`
		Language        string    `json:"language"`
		ForksCount      int       `json:"forks_count"`
		StargazersCount int       `json:"stargazers_count"`
		OpenIssuesCount int       `json:"open_issues_count"`
		WatchersCount   int       `json:"watchers_count"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}
	if err := json.Unmarshal(body, &apiRepo); err != nil {
		return nil, err
	}

	repo := &model.Repository{
		Name:            apiRepo.Name,
		Description:     apiRepo.Description,
		URL:             apiRepo.HTMLURL,
		Language:        apiRepo.Language,
		ForksCount:      apiRepo.ForksCount,
		StarsCount:      apiRepo.StargazersCount,
		OpenIssuesCount: apiRepo.OpenIssuesCount,
		WatchersCount:   apiRepo.WatchersCount,
		CreatedAt:       apiRepo.CreatedAt,
		UpdatedAt:       apiRepo.UpdatedAt,
	}

	return repo, nil
}
