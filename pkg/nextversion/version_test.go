package nextversion_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
)

func TestVersions_OpenRepository_Success(t *testing.T) {

	// Call the Versions function
	opts := &nextversion.Options{
		Repo:           "../../fixtures/no-valid-tags",
		DefaultCurrent: "v1.0.0",
		Prestable:      false,
	}
	result, err := nextversion.Versions(opts)
	if err != nil {
		t.Fatalf("Versions failed: %v", err)
	}

	assert.False(t, result.HasCurrentVersion)
	assert.Equal(t, opts.DefaultCurrent, result.CurrentVersion)
}
