package github

import (
	"commit-monitor/internal/model"
	"encoding/json"
	"fmt"
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
			SHA:            apiCommit.SHA,
			RepositoryName: "chromium",
		}
	}

	return commits, nil
}
