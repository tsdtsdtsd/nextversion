package nextversion

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type Tag struct {
	Name   string
	Commit object.Commit
	Semver *semver.Version
	exists bool
}

// GitRepository
type GitRepository interface {
	Tags() (storer.ReferenceIter, error)
	TagObject(h plumbing.Hash) (*object.Tag, error)
}

// LastTag returns an annotated tag object witch has no prerelease identifier and is the greates semantic version
func LastTag(repo GitRepository) (*Tag, error) {

	lastTag := Tag{
		Commit: object.Commit{},
		Semver: &semver.Version{},
	}

	if repo == nil {
		return &lastTag, fmt.Errorf("repository is nil")
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

		// Store tag if this is the first valid tag or it has a greater version than the one we stored before
		if lastTag.Commit.Hash.IsZero() || semantic.GreaterThan(lastTag.Semver) {
			lastTag.Commit = *commit
			lastTag.Name = obj.Name
			lastTag.Semver = semantic
			lastTag.exists = true
		}

		return nil
	})

	return &lastTag, err
}

func (tag *Tag) Exists() bool {
	return tag.exists
}
