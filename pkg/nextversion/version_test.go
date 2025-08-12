package nextversion_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
)

func TestVersions_OpenRepository_Success(t *testing.T) {

	opts := &nextversion.Options{
		Repo: "../../fixtures/no-valid-tags",
	}

	// Call the Versions function

	result, err := nextversion.Versions(opts)
	if err != nil {
		t.Fatalf("Versions failed: %v", err)
	}

	assert.False(t, result.HasCurrentVersion)
	assert.Equal(t, opts.DefaultCurrent, result.CurrentVersion)
}

func TestVersions_OpenRepository_RepositoryNotExists(t *testing.T) {

	// Directory does not exist

	opts := &nextversion.Options{
		Repo: "../../fixtures/does-not-exist",
	}

	// Call the Versions function
	_, err := nextversion.Versions(opts)

	assert.Error(t, err)
	assert.ErrorIs(t, err, git.ErrRepositoryNotExists)

	// Directory exists but is not a git repository

	opts = &nextversion.Options{
		Repo: "../../fixtures",
	}

	_, err = nextversion.Versions(opts)

	assert.Error(t, err)
	assert.ErrorIs(t, err, git.ErrRepositoryNotExists)
}

const (
	repoNoValidTags = "../../fixtures/no-valid-tags"
	repoValidTag    = "../../fixtures/valid-tag"
)

func TestVersions_NoValidTags_Defaults(t *testing.T) {

	opts := &nextversion.Options{
		Repo: repoNoValidTags,
	}

	// Call the Versions function
	result, err := nextversion.Versions(opts)

	assert.NoError(t, err)

	assert.Equal(t, "v0.0.0", result.CurrentVersion, "CurrentVersion mismatch")
	assert.Equal(t, "0.0.0", result.CurrentVersionStrict, "CurrentVersionStrict mismatch")
	assert.Equal(t, false, result.HasCurrentVersion, "HasCurrentVersion mismatch")
	assert.Equal(t, true, result.HasNextVersion, "HasNextVersion mismatch")
	assert.Equal(t, "v0.1.0", result.NextVersion, "NextVersion mismatch")
	assert.Equal(t, "0.1.0", result.NextVersionStrict, "NextVersionStrict mismatch")

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersion, "v0.1.0-rc+main.")
	}, "PrereleaseVersion mismatch")

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersionStrict, "0.1.0-rc+main.")
	}, "PrereleaseVersionStrict mismatch")
}

func TestVersions_NoValidTags_CustomDefaultCurrent(t *testing.T) {

	var (
		expectedCurrent       = "v3.1.2"
		expectedCurrentStrict = "3.1.2"
		expectedNext          = "v3.2.0"
		expectedNextStrict    = "3.2.0"
	)

	opts := &nextversion.Options{
		Repo:           repoNoValidTags,
		DefaultCurrent: expectedCurrent,
	}

	// Call the Versions function
	result, err := nextversion.Versions(opts)

	assert.NoError(t, err)

	assert.Equal(t, expectedCurrent, result.CurrentVersion, "CurrentVersion mismatch")
	assert.Equal(t, expectedCurrentStrict, result.CurrentVersionStrict, "CurrentVersionStrict mismatch")
	assert.Equal(t, false, result.HasCurrentVersion, "HasCurrentVersion mismatch")
	assert.Equal(t, true, result.HasNextVersion, "HasNextVersion mismatch")
	assert.Equal(t, expectedNext, result.NextVersion, "NextVersion mismatch")
	assert.Equal(t, expectedNextStrict, result.NextVersionStrict, "NextVersionStrict mismatch")

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersion, expectedNext+"-rc+main.")
	}, "PrereleaseVersion mismatch")

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersionStrict, expectedNextStrict+"-rc+main.")
	}, "PrereleaseVersionStrict mismatch")
}

func TestVersions_NoValidTags_ForceStable(t *testing.T) {

	var (
		expectedCurrent       = "v0.0.0"
		expectedCurrentStrict = "0.0.0"
		expectedNext          = "v1.0.0"
		expectedNextStrict    = "1.0.0"
	)

	opts := &nextversion.Options{
		Repo:        repoNoValidTags,
		ForceStable: true,
	}

	// Call the Versions function
	result, err := nextversion.Versions(opts)

	assert.NoError(t, err)

	assert.Equal(t, expectedCurrent, result.CurrentVersion, "CurrentVersion mismatch")
	assert.Equal(t, expectedCurrentStrict, result.CurrentVersionStrict, "CurrentVersionStrict mismatch")
	assert.Equal(t, false, result.HasCurrentVersion, "HasCurrentVersion mismatch")
	assert.Equal(t, true, result.HasNextVersion, "HasNextVersion mismatch")
	assert.Equal(t, expectedNext, result.NextVersion, "NextVersion mismatch")
	assert.Equal(t, expectedNextStrict, result.NextVersionStrict, "NextVersionStrict mismatch")

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersion, expectedNext+"-rc+main.")
	}, "PrereleaseVersion mismatch "+result.PrereleaseVersion)

	assert.Condition(t, func() bool {
		return strings.HasPrefix(result.PrereleaseVersionStrict, expectedNextStrict+"-rc+main.")
	}, "PrereleaseVersionStrict mismatch "+result.PrereleaseVersionStrict)
}

func TestSanitizeDockerTag(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"valid-tag_1.2.3",
			"valid-tag_1.2.3",
		},
		{
			"-invalid-start",
			fmt.Sprintf("%sinvalid-start", nextversion.DockerTagCharacterSubstitution),
		},
		{
			"hello world!",
			fmt.Sprintf("hello%sworld%s", nextversion.DockerTagCharacterSubstitution, nextversion.DockerTagCharacterSubstitution),
		},
		{
			"this/is\\not@valid*tag",
			fmt.Sprintf("this%sis%snot%svalid%stag",
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution),
		},
		{
			"", "",
		},
		{
			"@#$%",
			fmt.Sprintf("%s%s%s%s",
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution,
				nextversion.DockerTagCharacterSubstitution),
		},
		{
			"1valid", "1valid",
		},
		{
			string(make([]byte, 200)),
			strings.Repeat(nextversion.DockerTagCharacterSubstitution, 128),
		},
	}

	for _, test := range tests {
		output := nextversion.SanitizeDockerTag(test.input)
		assert.LessOrEqual(t, len(output), 128, "tag must not have more than 128 characters")
		assert.Equal(t, output, test.expected)
	}
}
