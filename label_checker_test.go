package main

import (
	"testing"
)

func TestLabelChecker_Check_Success(t *testing.T) {
	cases := []struct {
		labels []string
		expected bool
	}{
		{labels: []string{"test"}, expected: false},
		{labels: []string{"test1", "test2"}, expected: false},
		{labels: []string{"bug"}, expected: true},
		{labels: []string{"good first issue", "test1", "test2"}, expected: true},
	}

	checker, mux, _, tearDown := setupChecker()
	defer tearDown()

	number := 1
	setPullRequestHandler(mux, number, `{"labels":[{"name":"bug"},{"name":"good first issue"}]}`)

	for i, tc := range cases {
		result, err := checker.Check(number, tc.labels)
		if err != nil {
			t.Fatalf("#%d LabelChecker.Check returned unexpected error: %v", i, err)
		}

		if result != tc.expected {
			t.Errorf("#%d LabelChecker.Check returned %+v, want %+v", i, result, tc.expected)
		}
	}
}

func TestLabelChecker_Check_Fail(t *testing.T) {
	checker, mux, _, tearDown := setupChecker()
	defer tearDown()

	number := 1
	setPullRequestHandler(mux, number, `{"labels":[]}`)

	_, err := checker.Check(number, []string{"bug"})
	if err == nil {
		t.Fatalf("LabelChecker.Check is supposed to return error")
	}
}