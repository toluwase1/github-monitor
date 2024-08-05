package github

import (
	"commit_fetcher/internal/domain/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{Token: token}
}

func (c *Client) GetCommits(since, until time.Time) ([]model.Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/chromium/chromium/commits?since=%s&until=%s",
		since.Format(time.RFC3339), until.Format(time.RFC3339))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.Token)

	client := &http.Client{}
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
				Name string    `json:"name"`
				Date time.Time `json:"date"`
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
			SHA:     apiCommit.SHA,
			Message: apiCommit.Commit.Message,
			Author:  apiCommit.Commit.Author.Name,
			Date:    apiCommit.Commit.Author.Date,
			URL:     apiCommit.URL,
		}
	}

	return commits, nil
}
