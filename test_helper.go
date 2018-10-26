package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	testGitHubOwner = "shuheiktgw"
	testGitHubRepo  = "bump-reviewer"
	testGitHubToken = "abcdefg12345"
)

// setup sets up a test HTTP server along with a GitHubClient that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *GitHubClient, mux *http.ServeMux, serverURL string, tearDown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewGitHubClient(testGitHubOwner, testGitHubRepo, testGitHubToken)
	u, _ := url.Parse(server.URL + "/")
	client.Client.BaseURL = u

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func setupChecker() (reviewer *LabelChecker, mux *http.ServeMux, url string, tearDown func()) {
	client, mux, url, tearDown := setup()
	return &LabelChecker{client}, mux, url, tearDown
}

func setPullRequestHandler(mux *http.ServeMux, number int, pr string) {
	mux.HandleFunc(fmt.Sprintf("/repos/%v/%v/pulls/%d", testGitHubOwner, testGitHubRepo, number), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, pr)
	})
}