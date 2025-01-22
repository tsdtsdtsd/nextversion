package nextversion

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

const (
	hashLength          = 7
	prereleaseShorthand = "rc"

	defaultCurrent = "v0.0.0"
)

type Result struct {
	CurrentVersion          string `json:"current"`
	CurrentVersionStrict    string `json:"current-strict"`
	HasCurrentVersion       bool   `json:"has-current"`
	NextVersion             string `json:"next"`
	NextVersionStrict       string `json:"next-strict"`
	HasNextVersion          bool   `json:"has-next"`
	PrereleaseVersion       string `json:"prerelease"`
	PrereleaseVersionStrict string `json:"prerelease-strict"`
}

func Versions(opts *Options) (*Result, error) {

	var (
		err    error
		result = &Result{}
	)

	if opts.DefaultCurrent == "" {
		opts.DefaultCurrent = defaultCurrent
	}

	// Open repository
	repo, err := git.PlainOpen(opts.Repo)
	if err != nil {
		return result, fmt.Errorf("failed to open repository at '%s': %w", opts.Repo, err)
	}

	// Fetch last version tag details
	lastTag, err := LastTag(repo)
	if err != nil {
		return result, fmt.Errorf("failed to determine last version tag: %w", err)
	}

	// Current version info
	result.CurrentVersion = opts.DefaultCurrent
	if lastTag.Exists() {
		result.HasCurrentVersion = true
		result.CurrentVersion = lastTag.Semver.Original()
	}
	result.CurrentVersionStrict = strings.TrimPrefix(result.CurrentVersion, versionPrefix)

	bump, err := NewBumper(result.CurrentVersion, opts.PreStable, opts.ForceStable)
	if err != nil {
		return result, fmt.Errorf("failed to create version bumper: %w", err)
	}

	// Init log iterator
	logIterator, err := repo.Log(&git.LogOptions{
		Since: &lastTag.Commit.Committer.When,
	})
	if err != nil {
		return result, fmt.Errorf("failed to fetch git log entries: %w", err)
	}

	machine := parser.NewMachine(
		conventionalcommits.WithTypes(conventionalcommits.TypesConventional),
		conventionalcommits.WithBestEffort(),
	)

	// Iterate logs and collect changes
	logIterator.ForEach(func(c *object.Commit) error {

		// Skip commit of last tag
		if c.Hash == lastTag.Commit.Hash {
			return nil
		}

		message, err := machine.Parse([]byte(strings.TrimSpace(c.Message)))
		if err != nil {
			// Skip unparsable commit messages
			// githubactions.Warningf("skipping unparsable commit %s: %s", c.Hash.String(), err)
			return nil
		}

		bump.CollectChange(message.IsBreakingChange(), message.IsFeat(), message.IsFix())
		return nil
	})

	result.NextVersion = bump.Next()
	result.NextVersionStrict = strings.TrimPrefix(result.NextVersion, versionPrefix)

	result.HasNextVersion = bump.Next() != bump.Current()
	// Prerelease

	prereleaseSuffix, err := getPrereleaseSuffix(repo)
	if err != nil {
		return result, fmt.Errorf("failed to get prerelease suffix: %w", err)
	}

	result.PrereleaseVersion = result.NextVersion + prereleaseSuffix
	result.PrereleaseVersionStrict = result.NextVersionStrict + prereleaseSuffix

	return result, nil
}

func getPrereleaseSuffix(repo *git.Repository) (string, error) {

	// Fetch current HEAD
	head, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to fetch HEAD reference: %w", err)
	}

	var (
		reg        = regexp.MustCompile(`[^a-zA-Z0-9]+`)
		branchName = reg.ReplaceAllString(head.Name().Short(), "-")
		hash       = head.Hash().String()
	)

	if len(hash) != 40 {
		return "", errors.New("given hash has an invalid length")
	}

	result := fmt.Sprintf("-%s+%s.%s", prereleaseShorthand, branchName, hash[0:hashLength])

	return result, nil
}
