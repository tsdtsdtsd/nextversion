# nextversion

[![Basic CI](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tsdtsdtsd/nextversion/actions/workflows/ci.yml)

> Automatic semantic versioning utility

nextversion extracts version information from git tags of your repository.  
After analyzing the commits since that tag (based on [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0/)),
it will provide a logical new version number.
