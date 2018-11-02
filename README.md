github-label-checker
====
[![GitHub release](http://img.shields.io/github/release/shuheiktgw/github-label-checker.svg?style=flat-square)](release)
[![CircleCI](https://circleci.com/gh/shuheiktgw/github-label-checker.svg?style=svg)](https://circleci.com/gh/shuheiktgw/github-label-checker)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

`github-label-checker` is a CLI tool to check if the specified PR has the given labels or not.

## Example

```
github-label-checker \ 
 -owner shuheiktgw \
 -repo bump-reviewer \
 -token abcdefg \ 
 -number 10 \
 bumps, bug
```

When the PR has the specified label `github-label-checker` exits with 0, otherwise exits with non-zero value. Please be aware that `github-label-checker` exits with 0 if at least one label matched with the upstream labels.

## Usage

```
github-label-checker [options...] LABELS

OPTIONS:
  --owner value, -o value   specifies GitHub Owner
  --repo value, -r value    specifies GitHub Repository Name
  --token value, -v value   specifies GitHub Personal Access Token
  --number value, -n value  specifies GitHub Pull Request Number to review
  --regex                   compares labels using regular expressions
  --version, -v             prints the current version
  --help, -h                prints help
```

## Author
[Shuhei Kitagawa](https://github.com/shuheiktgw)