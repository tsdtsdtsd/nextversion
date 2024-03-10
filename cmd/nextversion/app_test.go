package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

// Returns the app instance with the correct minimal value set
func TestNewAppReturnsCorrectInstance(t *testing.T) {
	app := newApp()

	assert.Equal(t, "nextversion", app.Name)
	assert.Equal(t, version, app.Version)
	assert.NotEmpty(t, app.Flags)
}

// appAction fails due to missing repo
func TestAppActionError(t *testing.T) {
	// Set up test case
	ctx := cli.NewContext(nil, nil, nil)
	ctx.Set("repo", "") // Missing repo
	ctx.Set("format", "simple")
	ctx.Set("defaultCurrent", "1.0.0")

	// Call the function under test
	err := appAction(ctx)

	// Check the result
	assert.Error(t, err)
	// Additional assertions to check the error message or other details
}

// None of the flags have empty names or string represented values
func TestAppFlags_NoEmptyValuesOrAliases(t *testing.T) {
	flags := appFlags()

	for _, flag := range flags {
		assert.NotEmpty(t, flag.Names())
		assert.NotEmpty(t, flag.String())
	}
}

// Verify that a valid format value returns no error.
func TestValidFormatValue(t *testing.T) {
	ctx := &cli.Context{}
	values := []string{"simple", "json"}

	for _, value := range values {
		t.Run("test format '"+value+"'", func(t *testing.T) {
			err := verifyFormatValue(ctx, value)

			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}

// Verify that an empty format value returns an error.
func TestEmptyFormatValue(t *testing.T) {
	ctx := &cli.Context{}
	value := ""

	err := verifyFormatValue(ctx, value)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
