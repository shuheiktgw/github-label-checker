package main

import (
	"fmt"
	"regexp"
)

// LabelChecker checks if the labels are as expected
type LabelChecker struct {
	*GitHubClient
}

// Check checks if the specified PR has given labels
func(lc *LabelChecker) Check(number int, labels []string, regex bool) (bool, []string, error) {
	pr, err := lc.GetPullRequest(number)
	if err != nil {
		return false, nil, err
	}

	if len(pr.Labels) == 0 {
		return false, nil, fmt.Errorf("the specified pull request does not have any labels")
	}

	var upstreamLabels []string
	for _, lb := range pr.Labels {
		upstreamLabels = append(upstreamLabels, *lb.Name)
	}

	result, err := check(upstreamLabels, labels, regex)
	return result, upstreamLabels, err
}

func check(upstreamLabels []string, labels []string, regex bool) (bool, error) {
	for _, ul := range upstreamLabels {
		for _, l := range labels {
			if regex {
				r, err := regexp.Compile(l)
				if err != nil {
					return false, err
				}

				if r.Match([]byte(ul)) {
					return true, nil
				}
			}

			if ul == l {
				return true, nil
			}
		}
	}

	return false, nil
}



