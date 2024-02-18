package nextversion

import (
	"github.com/Masterminds/semver/v3"
)

const (
	VersionPrefix = "v"
)

type Bumper struct {
	current     *semver.Version
	preStable   bool
	hasBreaking bool
	hasFeature  bool
	hasFix      bool
}

func NewBumper(currentVersion string, preStable bool) (*Bumper, error) {
	v, err := semver.NewVersion(currentVersion)
	if err != nil {
		return nil, err
	}

	return &Bumper{
		current:   v,
		preStable: preStable,
	}, nil
}

func (b *Bumper) CollectChange(isBreaking, isFeature, isFix bool) {
	if isBreaking {
		b.hasBreaking = true
	}
	if isFeature {
		b.hasFeature = true
	}
	if isFix {
		b.hasFix = true
	}
}

func (b *Bumper) Next() string {

	switch b.preStable {
	case true:

		if b.hasBreaking && b.current.Major() > 0 {
			return VersionPrefix + b.current.IncMajor().String()
		}

		if b.hasBreaking || b.hasFeature {
			return VersionPrefix + b.current.IncMinor().String()
		}

	case false:

		if b.hasBreaking {
			return VersionPrefix + b.current.IncMajor().String()
		}

		if b.hasFeature {
			return VersionPrefix + b.current.IncMinor().String()
		}

	}

	if b.hasFix {
		return VersionPrefix + b.current.IncPatch().String()
	}

	return VersionPrefix + b.current.String()
}

func (b *Bumper) Current() string {
	return VersionPrefix + b.current.String()
}
