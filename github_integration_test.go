// +build integration

package main

import (
	"testing"
)

func TestGitHubClient_Integration_GetPullRequest(t *testing.T) {
	pr, err := integrationGitHubClient.GetPullRequest(1)

	if err != nil {
		t.Fatalf("GitHubClient.GetPullRequest returns unexpected error: %s", err)
	}

	if got, want := *pr.Labels[0].Name, "good first issue"; got != want {
		t.Fatalf("GitHubClient.GetPullRequest returns unexpected file: want: %s, got: %s", want, got)
	}
}