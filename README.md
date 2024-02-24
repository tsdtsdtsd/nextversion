# nextversion

> Automatic semantic versioning utility

![Latest Release Version][shields-version-img]
[![Godoc][godoc-image]][godoc-url]
![Build Status](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml/badge.svg)
[![Go Report Card][grc-image]][grc-url]
[![codecov][codecov-image]][codecov-url]
[![CodeQL](https://github.com/tsdtsdtsd/nextversion/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/tsdtsdtsd/nextversion/actions/workflows/codeql-analysis.yml)

<!-- Markdown link & img dfn's -->
[shields-version-img]: https://img.shields.io/github/v/release/tsdtsdtsd/nextversion
[godoc-image]: https://pkg.go.dev/badge/github.com/tsdtsdtsd/nextversion.svg
[godoc-url]: https://pkg.go.dev/github.com/tsdtsdtsd/nextversion/pkg/nextversion/
[grc-image]: https://goreportcard.com/badge/github.com/tsdtsdtsd/nextversion
[grc-url]: https://goreportcard.com/report/github.com/tsdtsdtsd/nextversion
[codecov-image]: https://codecov.io/gh/tsdtsdtsd/nextversion/branch/main/graph/badge.svg
[codecov-url]: https://codecov.io/gh/tsdtsdtsd/nextversion/tree/main

nextversion detects current application version based on git tags and suggests a bumped version based on [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0/).

## Constraints

- nextversion needs annotated git tags to determine the current app version
- nextversion does not create tags for you, it only generates a suggestion
- nextversion sticks to [Semantic Versioning](https://semver.org/) and [Conventional Commits](https://www.conventionalcommits.org):
  - your tags must follow the semver pattern without prerelease suffix and an optional `v` prefix (e.g. `v1.2.3` or `1.2.3`).
  - you have to make sure that commit messages follow the [Conventional Commits spec](https://www.conventionalcommits.org/en/v1.0.0/).  
    The used parser is in [best effort mode](https://github.com/leodido/go-conventionalcommits#best-effort)

## Usage

```sh
$ nextversion --help
NAME:
   nextversion - versioning helper tool

USAGE:
   nextversion [global options] command [command options] 

VERSION:
   v0.0.0-dev

DESCRIPTION:
   nextversion detects application version based on git tags and suggests a bumped version based on conventional commits.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --repo PATH, -r PATH                  PATH to a git repository (default: "./")
   --format FORMAT, -f FORMAT            Output FORMAT (simple, json) (default: "simple")
   --defaultCurrent VERSION, -d VERSION  Fallback current VERSION if none could be detected (default: "v0.0.0")
   --prestable, -p                       Pre-stable mode (default: false)
   --help, -h                            show help
   --version, -v                         print version of this tool (default: false)
```
