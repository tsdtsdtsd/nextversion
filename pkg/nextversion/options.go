package nextversion

// Options define some
type Options struct {
	// Repo is the repositories location
	Repo string

	// Format sets the output format
	Format string

	// DefaultCurrent will be used as fallback if a current version could not be determined via tags
	DefaultCurrent string

	// PreStable marks this repo as still in development
	// In this state, major changes will not increment major version if current version is pre 1.0.0
	PreStable bool
}
