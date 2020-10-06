package version

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/release"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Check for an update for the CLI
func CheckForUpdate() {
	s := spinner.New(spinner.CharSets[13], 30*time.Millisecond)
	s.Suffix = " Checking for update"
	s.Start()

	const repoURL = "https://github.com/Matt-Gleich/nuke"
	isOutdated, version, err := release.Check("v5.1.0", repoURL)
	if err != nil {
		statuser.Error("Failed to get current version number", err, 1)
	}
	s.Stop()
	if isOutdated {
		fmt.Println("\nPLEASE UPDATE TO NEW VERSION", version)
		fmt.Println("Get update from", repoURL, "or with your package manager")
		fmt.Println()
	}
}
