package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-github/github"
)

func TestGitHubClient_ListPullRequestsFiles(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	number := 1

	mux.HandleFunc(fmt.Sprintf("/repos/%v/%v/pulls/%d", testGitHubOwner, testGitHubRepo, number), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"labels": [{"name":"bug"}]}`)
	})

	cc, err := client.GetPullRequest(number)
	if err != nil {
		t.Fatalf("GitHubClient.GetPullRequest returned unexpected error: %v", err)
	}

	want := &github.PullRequest{Labels:[]*github.Label{{Name: github.String("bug")}}}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("GitHubClient.GetPullRequest returned %+v, want %+v", cc, want)
	}
}