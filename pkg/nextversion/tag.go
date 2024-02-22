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

	tag := Tag{
		Commit: object.Commit{},
		Semver: &semver.Version{},
	}

	tagRefs, err := repo.Tags()
	if err != nil {
		return &tag, fmt.Errorf("failed to fetch tag references: %w", err)
	}

	err = tagRefs.ForEach(func(ref *plumbing.Reference) error {

		obj, err := repo.TagObject(ref.Hash())
		if err != nil {
			// case plumbing.ErrObjectNotFound can be silently skipped
			return nil
		}

		// Skip non-semver tag names
		semantic, err := semver.NewVersion(obj.Name)
		if err != nil {
			// fmt.Println("non-semver", obj.Name)
			return nil
		}

		// Ignore prereleases
		if semantic.Prerelease() != "" {
			return nil
		}

		commit, err := obj.Commit()
		if err != nil {
			// TODO: should we notify the user?
			// fmt.Println("no commit")
			return err
		}

		// At this point we know that this is a valid tag object

		// Store tag and commit if this is the first iteration
		if tag.Commit.Hash.IsZero() {
			tag.Commit = *commit
			tag.Name = obj.Name
			tag.Semver = semantic
			tag.exists = true
			return nil
		}

		// Store tag if it's a greater version than the one we already stored
		if semantic.GreaterThan(tag.Semver) {
			tag.Commit = *commit
			tag.Name = obj.Name
			tag.Semver = semantic
			tag.exists = true
		}

		return nil

	})

	return &tag, err
}

func (v *Tag) Exists() bool {
	return v.exists
}
