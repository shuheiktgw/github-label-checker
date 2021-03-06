package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tcnksm/go-latest"
)

// The name of this application
const Name = "github-label-checker"

// The current version of bump-reviewer
const Version = "0.0.1"

// The owner of bump-reviewer
const Owner = "shuheiktgw"

// OutputVersion outputs current version of bump-reviewer. It also checks
// the latest release and adds a warning to update bump-reviewer
func OutputVersion() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s current version v%s\n", Name, Version)

	// Get the latest release
	verCheckCh := make(chan *latest.CheckResponse)
	go func() {
		githubTag := &latest.GithubTag{
			Owner:      Owner,
			Repository: Name,
		}

		res, err := latest.Check(githubTag, Version)

		// Ignore the error
		if err != nil {
			return
		}

		verCheckCh <- res
	}()

	select {
	case <-time.After(2 * time.Second):
	case res := <-verCheckCh:
		if res.Outdated {
			fmt.Fprintf(&b, "The latest version is v%s, please update\n", res.Current)
		}
	}

	return b.String()
}
