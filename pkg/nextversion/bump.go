package nextversion

import (
	"github.com/Masterminds/semver/v3"
)

const (
	versionPrefix = "v"
	stableVersion = "1.0.0"
)

type Bumper struct {
	current     *semver.Version
	preStable   bool
	forceStable bool
	hasBreaking bool
	hasFeature  bool
	hasFix      bool
}

func NewBumper(currentVersion string, preStable bool, forceStable bool) (*Bumper, error) {
	v, err := semver.NewVersion(currentVersion)
	if err != nil {
		return nil, err
	}

	return &Bumper{
		current:     v,
		preStable:   preStable,
		forceStable: forceStable,
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

	if b.forceStable && b.current.LessThan(semver.MustParse(stableVersion)) {
		return versionPrefix + stableVersion
	}

	switch b.preStable {
	case true:

		if b.hasBreaking && b.current.Major() > 0 {
			return versionPrefix + b.current.IncMajor().String()
		}

		if b.hasBreaking || b.hasFeature {
			return versionPrefix + b.current.IncMinor().String()
		}

	case false:

		if b.hasBreaking {
			return versionPrefix + b.current.IncMajor().String()
		}

		if b.hasFeature {
			return versionPrefix + b.current.IncMinor().String()
		}

	}

	if b.hasFix {
		return versionPrefix + b.current.IncPatch().String()
	}

	return versionPrefix + b.current.String()
}

func (b *Bumper) Current() string {
	return versionPrefix + b.current.String()
}
