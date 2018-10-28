package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	cases := []struct {
		command           string
		expectedOutStream string
		expectedErrStream string
		expectedExitCode  int
	}{
		{
			command:           "github-label-checker",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up github-label-checker: GitHub owner is missing\nPlease set it via `-o` option\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "github-label-checker -o shuheiktgw",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up github-label-checker: GitHub repository is missing\nPlease set it via `-r` option\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "github-label-checker -o shuheiktgw -r github-label-checker",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up github-label-checker: GitHub Personal Access Token is missing\nPlease set it via `-t` option\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "github-label-checker -o shuheiktgw -r github-label-checker -t 1234abcd",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up github-label-checker: Pull Request number is missing\nPlease set it via `-n` option\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "github-label-checker -o shuheiktgw -r github-label-checker -t 1234abcd -n 1",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up github-label-checker: Invalid arguments\nPlease specify at least one label\n\n",
			expectedExitCode: ExitCodeBadArgs,
		},
		{
			command:           "github-label-checker -o shuheiktgw -r github-label-checker -t 1234abcd -n 1 bug",
			expectedOutStream: "",
			expectedErrStream: "github-label-checker failed to review because of the following error.\n\n" +
				"GET https://api.github.com/repos/shuheiktgw/github-label-checker/pulls/1: 401 Bad credentials []\n\n" +
				"You might encounter a bug with github-label-checker, and if so, please report it to https://github.com/shuheiktgw/github-label-checker/issues\n\n",
			expectedExitCode: ExitCodeError,
		},
		{
			command:           "github-label-checker -v",
			expectedOutStream: fmt.Sprintf("github-label-checker current version v%s\n", Version),
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
	}

	for i, tc := range cases {
		outStream := new(bytes.Buffer)
		errStream := new(bytes.Buffer)

		cli := CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(tc.command, " ")

		if got := cli.Run(args); got != tc.expectedExitCode {
			t.Fatalf("#%d %q exits with %d, want %d", i, tc.command, got, tc.expectedExitCode)
		}

		if got := outStream.String(); got != tc.expectedOutStream {
			t.Fatalf("#%d Unexpected outStream has returned: want: %s, got: %s", i, tc.expectedOutStream, got)
		}

		if got := errStream.String(); got != tc.expectedErrStream {
			t.Fatalf("#%d Unexpected errStream has returned: want: %s, got: %s", i, tc.expectedErrStream, got)
		}
	}
}
