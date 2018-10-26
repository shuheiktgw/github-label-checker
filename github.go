package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GitHubClient is a clint to interact with Github API
type GitHubClient struct {
	Owner, Repo string
	Client      *github.Client
}

// NewGitHubClient creates and initializes a new GitHubClient
func NewGitHubClient(owner, repo, token string) *GitHubClient {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	return &GitHubClient{
		Owner:  owner,
		Repo:   repo,
		Client: client,
	}
}

// GetPullRequest get a single PR
func (c *GitHubClient) GetPullRequest(number int) (*github.PullRequest, error) {
	pr, res, err := c.Client.PullRequests.Get(context.TODO(), c.Owner, c.Repo, number)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PullRequests.Get returns invalid status: %s", res.Status)
	}

	return pr, nil
}