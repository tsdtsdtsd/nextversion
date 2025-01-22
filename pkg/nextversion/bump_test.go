package nextversion_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
)

type bumperTestcase struct {
	OldVersion      string
	PreStable       bool
	ForceStable     bool
	ChangeMajor     bool
	ChangeMinor     bool
	ChangePatch     bool
	ExpectedVersion string
}

func TestCurrent(t *testing.T) {

	oldVersion := "v1.0.1"

	// Pre-stable, force stable
	b, err := nextversion.NewBumper(oldVersion, true, true)
	v := b.Current()

	assert.NoError(t, err)
	assert.Equal(t, oldVersion, v)

	// Pre-stable, don't force stable
	b, err = nextversion.NewBumper(oldVersion, true, false)
	v = b.Current()

	assert.NoError(t, err)
	assert.Equal(t, oldVersion, v)

	// Not pre-stable, force stable
	b, err = nextversion.NewBumper(oldVersion, false, true)
	v = b.Current()

	assert.NoError(t, err)
	assert.Equal(t, oldVersion, v)

	// Not pre-stable, don't force stable
	b, err = nextversion.NewBumper(oldVersion, false, false)
	v = b.Current()

	assert.NoError(t, err)
	assert.Equal(t, oldVersion, v)
}

func TestPrestable(t *testing.T) {
	testTable := []bumperTestcase{
		// ChangeMajor = true
		{
			OldVersion:      "1.1.0",
			PreStable:       true,
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v2.0.0",
		},

		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v0.2.0",
		},

		// ChangeMajor = false
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v0.1.1",
		},
		{
			OldVersion:      "0.1.0",
			PreStable:       true,
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v0.1.0",
		},
	}

	runValidBumperChecksOnTable(t, testTable)
}

func TestNotPrestable(t *testing.T) {
	testTable := []bumperTestcase{
		// ChangeMajor = true
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v1.0.0",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v1.0.0",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v1.0.0",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v1.0.0",
		},

		{
			OldVersion:      "1.0.0",
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v2.0.0",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     true,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v2.0.0",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v2.0.0",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     true,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v2.0.0",
		},

		// ChangeMajor = false
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v0.2.0",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v0.1.1",
		},
		{
			OldVersion:      "0.1.0",
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v0.1.0",
		},

		{
			OldVersion:      "1.0.0",
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     true,
			ExpectedVersion: "v1.1.0",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     false,
			ChangeMinor:     true,
			ChangePatch:     false,
			ExpectedVersion: "v1.1.0",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     true,
			ExpectedVersion: "v1.0.1",
		},
		{
			OldVersion:      "1.0.0",
			ChangeMajor:     false,
			ChangeMinor:     false,
			ChangePatch:     false,
			ExpectedVersion: "v1.0.0",
		},
	}

	runValidBumperChecksOnTable(t, testTable)
}

// func TestForceStable(t *testing.T) {
// 	testTable := []bumperTestcase{
// 		// ChangeMajor = true
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     true,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v1.0.0",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     true,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v1.0.0",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     false,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v1.0.0",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     false,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v1.0.0",
// 		},

// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     true,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v2.0.0",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     true,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v2.0.0",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     false,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v2.0.0",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     true,
// 			ChangeMinor:     false,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v2.0.0",
// 		},

// 		// ChangeMajor = false
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     true,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v0.2.0",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     true,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v0.2.0",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     false,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v0.1.1",
// 		},
// 		{
// 			OldVersion:      "0.1.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     false,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v0.1.0",
// 		},

// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     true,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v1.1.0",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     true,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v1.1.0",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     false,
// 			ChangePatch:     true,
// 			ExpectedVersion: "v1.0.1",
// 		},
// 		{
// 			OldVersion:      "1.0.0",
// 			ForceStable:     true,
// 			ChangeMajor:     false,
// 			ChangeMinor:     false,
// 			ChangePatch:     false,
// 			ExpectedVersion: "v1.0.0",
// 		},
// 	}

// 	runValidBumperChecksOnTable(t, testTable)
// }

func TestInvalidOldVersion(t *testing.T) {

	testTable := []bumperTestcase{
		{
			OldVersion: "",
		},
		{
			OldVersion: "abc",
		},
		{
			OldVersion: "v1.0.0.0",
		},
		{
			OldVersion: "v.",
		},
		{
			OldVersion: "v1..",
		},
	}

	for _, testCase := range testTable {
		_, err := nextversion.NewBumper(testCase.OldVersion, testCase.PreStable, false)

		assert.Error(t, err)
	}
}

func runValidBumperChecksOnTable(t *testing.T, testTable []bumperTestcase) {
	for _, testCase := range testTable {
		b, err := nextversion.NewBumper(testCase.OldVersion, testCase.PreStable, false)

		b.CollectChange(testCase.ChangeMajor, testCase.ChangeMinor, testCase.ChangePatch)
		v := b.Next()

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedVersion, v)
	}
}
