package nextversion

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Tag struct {
	Name   string
	Commit object.Commit
	Semver *semver.Version
	exists bool
}

// LastTag returns an annotated tag object witch has no prerelease identifier and is the greates semantic version
func LastTag(repo *git.Repository) (*Tag, error) {

	lastTag := Tag{
		Commit: object.Commit{},
		Semver: &semver.Version{},
	}

	tagRefs, err := repo.Tags()
	if err != nil {
		return &lastTag, fmt.Errorf("failed to fetch tag references: %w", err)
	}

	err = tagRefs.ForEach(func(ref *plumbing.Reference) error {

		// Skip if this is not an annotated tag
		obj, err := repo.TagObject(ref.Hash())
		if err != nil {
			return nil
		}

		// Skip non-semver tag names
		semantic, err := semver.NewVersion(obj.Name)
		if err != nil {
			return nil
		}

		// Ignore prereleases
		if semantic.Prerelease() != "" {
			return nil
		}

		commit, err := obj.Commit()
		if err != nil {
			// TODO: should we notify the user?
			return err
		}

		// At this point we know that this is a valid tag object

		// Store tag and commit if this is the first iteration
		if lastTag.Commit.Hash.IsZero() {
			lastTag.Commit = *commit
			lastTag.Name = obj.Name
			lastTag.Semver = semantic
			lastTag.exists = true
			return nil
		}

		// Store tag if it's a greater version than the one we already stored
		if semantic.GreaterThan(lastTag.Semver) {
			lastTag.Commit = *commit
			lastTag.Name = obj.Name
			lastTag.Semver = semantic
			lastTag.exists = true
		}

		return nil

	})

	return &lastTag, err
}

func (v *Tag) Exists() bool {
	return v.exists
}
