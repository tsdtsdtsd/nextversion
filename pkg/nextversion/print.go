package nextversion

import (
	"encoding/json"
	"fmt"
)

func Print(versions *Result, format string) error {

	switch format {
	case "shell":
		// TODO: OS specific EOL?
		fmt.Printf("CURRENT=%s\n", versions.CurrentVersion)
		fmt.Printf("CURRENT_STRICT=%s\n", versions.CurrentVersionStrict)
		fmt.Printf("HAS_CURRENT=%t\n", versions.HasCurrentVersion)
		fmt.Printf("NEXT=%s\n", versions.NextVersion)
		fmt.Printf("NEXT_STRICT=%s\n", versions.NextVersionStrict)
		fmt.Printf("HAS_NEXT=%t\n", versions.HasNextVersion)
		fmt.Printf("PRERELEASE=%s\n", versions.PrereleaseVersion)
		fmt.Printf("PRERELEASE_STRICT=%s\n", versions.PrereleaseVersionStrict)

	case "json":
		jsonString, err := json.Marshal(versions)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		fmt.Printf("%s", jsonString)
	}

	return nil
}
