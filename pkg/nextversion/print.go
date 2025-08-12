package nextversion

import (
	"encoding/json"
	"fmt"
	"io"
)

func Print(writer io.Writer, versions *Result, format string) error {

	switch format {
	case "simple":

		// TODO: OS specific EOL?

		fmt.Fprintf(writer, "CURRENT=%s\n", versions.CurrentVersion)
		fmt.Fprintf(writer, "CURRENT_STRICT=%s\n", versions.CurrentVersionStrict)
		fmt.Fprintf(writer, "HAS_CURRENT=%t\n", versions.HasCurrentVersion)
		fmt.Fprintf(writer, "NEXT=%s\n", versions.NextVersion)
		fmt.Fprintf(writer, "NEXT_STRICT=%s\n", versions.NextVersionStrict)
		fmt.Fprintf(writer, "HAS_NEXT=%t\n", versions.HasNextVersion)
		fmt.Fprintf(writer, "PRERELEASE=%s\n", versions.PrereleaseVersion)
		fmt.Fprintf(writer, "PRERELEASE_STRICT=%s\n", versions.PrereleaseVersionStrict)
		fmt.Fprintf(writer, "PRERELEASE_DOCKER_TAG=%s\n", versions.PrereleaseDockerTagVersion)

	case "json":

		outJSON, err := json.Marshal(versions)
		if err != nil {
			return fmt.Errorf("failed to marshal to JSON: %w", err)
		}

		writer.Write(outJSON)
	}

	return nil
}
