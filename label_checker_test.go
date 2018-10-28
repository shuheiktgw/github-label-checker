package main

import (
	"testing"
)

func TestLabelChecker_Check_Success(t *testing.T) {
	cases := []struct {
		labels []string
		casual bool
		expected bool
	}{
		{labels: []string{"test"}, casual: true, expected: false},
		{labels: []string{"test1", "test2"}, casual: true, expected: false},
		{labels: []string{"bug"}, casual: true, expected: true},
		{labels: []string{"bu"}, casual: true, expected: true},
		{labels: []string{"bu"}, casual: false, expected: false},
		{labels: []string{"bug"}, casual: false, expected: true},
		{labels: []string{"good first issue", "test1", "test2"}, casual: true, expected: true},
	}

	checker, mux, _, tearDown := setupChecker()
	defer tearDown()

	number := 1
	setPullRequestHandler(mux, number, `{"labels":[{"name":"bug"},{"name":"good first issue"}]}`)

	for i, tc := range cases {
		result, _, err := checker.Check(number, tc.labels, tc.casual)
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

	_, _, err := checker.Check(number, []string{"bug"}, false)
	if err == nil {
		t.Fatalf("LabelChecker.Check is supposed to return error")
	}
}