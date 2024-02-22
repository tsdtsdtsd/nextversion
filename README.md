# nextversion

[![Basic CI](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml)

> Automatic semantic versioning utility

nextversion detects current application version based on git tags and suggests a bumped version based on [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0/).

## Constraints

- nextversion needs annotated git tags to determine the current app version.  
  They must follow [semver](https://semver.org/) pattern without prerelease prefix and an optional `v` prefix (e.g. `v1.2.3` or `1.2.3`).
- nextversion does not create tags for you, it only generates a suggestion.
- nextversion sticks to _Semantic Versioning_ and _Conventional Commits_.  
  You have to make sure that commit messages follow the [Conventional Commits specifications](https://www.conventionalcommits.org/en/v1.0.0/),
  otherwise you will not get meaningful results from nextversion.

