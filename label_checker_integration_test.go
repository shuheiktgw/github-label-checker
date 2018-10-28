// +build integration

package main

import (
	"testing"
	"time"
)

func TestLabelChecker_Integration_Check_Success(t *testing.T) {
	cases := []struct {
		prNum int
		labels []string
		casual bool
		expected bool
	}{
		{prNum: 1, labels: []string{"test"}, casual: true, expected: false},
		{prNum: 1, labels: []string{"test1", "test2"}, casual: true, expected: false},
		{prNum: 1, labels: []string{"good first"}, casual: true, expected: true},
		{prNum: 1, labels: []string{"good first"}, casual: false, expected: false},
		{prNum: 1, labels: []string{"good first issue"}, casual: false, expected: true},
		{prNum: 1, labels: []string{"good first issue", "test1", "test2"}, expected: true},
	}

	for i, tc := range cases {
		time.Sleep(2 * time.Second)

		lc := LabelChecker{integrationGitHubClient}
		result, _, err := lc.Check(tc.prNum, tc.labels, tc.casual)

		if err != nil {
			t.Fatalf("#%d Unexpected error occurred from LabelChecker.Check: %s", i, err)
		}

		if result != tc.expected {
			t.Fatalf("#%d Unexpected value has returned from LabelChecker.Check: want: %v, got: %v", i, tc.expected, result)
		}
	}
}

func TestLabelChecker_Integration_Check_Fail(t *testing.T) {
	cases := []struct {
		prNum int
	}{
		{prNum: 100},
		{prNum: 4},
	}

	for i, tc := range cases {
		time.Sleep(2 * time.Second)

		lc := LabelChecker{integrationGitHubClient}
		_, _, err := lc.Check(tc.prNum, []string{"test"}, true)

		if err == nil {
			t.Fatalf("#%d LabelChecker.Check is supposed to return error", i)
		}
	}
}