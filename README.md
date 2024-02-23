# nextversion

[![Basic CI](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml)

> Automatic semantic versioning utility

nextversion detects current application version based on git tags and suggests a bumped version based on [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0/).

## Constraints

- nextversion needs annotated git tags to determine the current app version
- nextversion does not create tags for you, it only generates a suggestion
- nextversion sticks to [Semantic Versioning](https://semver.org/) and [Conventional Commits](https://www.conventionalcommits.org):
  - your tags must follow the semver pattern without prerelease suffix and an optional `v` prefix (e.g. `v1.2.3` or `1.2.3`).
  - you have to make sure that commit messages follow the [Conventional Commits spec](https://www.conventionalcommits.org/en/v1.0.0/).  
    The used parser is in [best effort mode](https://github.com/leodido/go-conventionalcommits#best-effort)

