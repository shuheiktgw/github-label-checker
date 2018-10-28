package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeLabelUnmatched
	ExitCodeError
	ExitCodeParseFlagsError
	ExitCodeBadArgs
	ExitCodeInvalidFlagError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		owner   string
		repo    string
		token   string
		number  int
		regex bool
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, usage)
	}

	flags.StringVar(&owner, "owner", "", "")
	flags.StringVar(&owner, "o", "", "")

	flags.StringVar(&repo, "repo", "", "")
	flags.StringVar(&repo, "r", "", "")

	flags.StringVar(&token, "token", "", "")
	flags.StringVar(&token, "t", "", "")

	flags.IntVar(&number, "number", 0, "")
	flags.IntVar(&number, "n", 0, "")

	flags.BoolVar(&regex, "regex", false, "")

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if version {
		fmt.Fprint(cli.outStream, OutputVersion())
		return ExitCodeOK
	}

	if len(owner) == 0 {
		fmt.Fprintf(cli.errStream, "Failed to set up github-label-checker: GitHub owner is missing\n"+
			"Please set it via `-o` option\n\n")
		return ExitCodeInvalidFlagError
	}

	if len(repo) == 0 {
		fmt.Fprintf(cli.errStream, "Failed to set up github-label-checker: GitHub repository is missing\n"+
			"Please set it via `-r` option\n\n")
		return ExitCodeInvalidFlagError
	}

	if len(token) == 0 {
		fmt.Fprintf(cli.errStream, "Failed to set up github-label-checker: GitHub Personal Access Token is missing\n"+
			"Please set it via `-t` option\n\n")
		return ExitCodeInvalidFlagError
	}

	if number == 0 {
		fmt.Fprintf(cli.errStream, "Failed to set up github-label-checker: Pull Request number is missing\n"+
			"Please set it via `-n` option\n\n")
		return ExitCodeInvalidFlagError
	}

	labels := flags.Args()
	if len(labels) < 1 {
		fmt.Fprintf(cli.errStream, "Failed to set up github-label-checker: Invalid arguments\n"+
			"Please specify at least one label\n\n")
		return ExitCodeBadArgs
	}

	client := NewGitHubClient(owner, repo, token)
	reviewer := LabelChecker{client}

	ok, upstreams, err := reviewer.Check(number, labels, regex)
	if err != nil {
		fmt.Fprintf(cli.errStream, `github-label-checker failed to review because of the following error.

%s

You might encounter a bug with github-label-checker, and if so, please report it to https://github.com/shuheiktgw/github-label-checker/issues

`, err)
		return ExitCodeError
	}

	if !ok {
		fmt.Fprintf(cli.errStream, "The pull request's labels did not match with the specified ones.\n" +
			"Upstream labels: %v\n" +
			"Specified labels: %v\n\n", upstreams, labels)
		return ExitCodeLabelUnmatched
	}

	fmt.Fprintf(cli.outStream, "Success. The pull request's labels matched with the specified ones\n\n")
	return ExitCodeOK
}

var usage = `Usage: github-label-checker [options...] LABELS

github-label-checker is a command to check the pull request's labels

OPTIONS:
  --owner value, -o value   specifies GitHub Owner
  --repo value, -r value    specifies GitHub Repository Name
  --token value, -v value   specifies GitHub Personal Access Token
  --number value, -n value  specifies GitHub Pull Request Number to review
  --regex                   compares labels using regular expressions
  --version, -v             prints the current version
  --help, -h                prints help

`
