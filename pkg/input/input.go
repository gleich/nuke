package input

import (
	"os"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/statuser/v2"
)

// Ask the user what terminal they are executing from
func ExecutingTerm(runningApps, ignoredApps []string, ignoredRunningFrom bool) []string {
	cleanedApps := []string{}

	// Removing ignored applications
	for _, app := range runningApps {
		var isIgnored bool
		for _, ignoredApp := range ignoredApps {
			if strings.Trim(app, "\n") == strings.Trim(ignoredApp, "\n") {
				isIgnored = true
			}
		}
		if !isIgnored {
			cleanedApps = append(cleanedApps, app)
		}
	}
	sort.Strings(cleanedApps)

	// Automatically select terminal as kitty if $TERM is xterm-kitty
	var program string

	if !ignoredRunningFrom {
		for _, app := range runningApps {
			if app == "kitty" && os.Getenv("TERM") == "xterm-kitty" {
				statuser.Success("Kitty terminal automatically detected")
				program = "kitty"
			}
		}

		if program == "" {
			err := survey.AskOne(
				&survey.Select{
					Message:  "Running nuke from (what app you're currently using)",
					Options:  cleanedApps,
					PageSize: 25,
				},
				&program)
			if err != nil {
				statuser.Error("Failed to get running terminal", err, 1)
			}
		}
	}

	var foundRunning bool
	cleanedApps2 := []string{}
	for _, app := range cleanedApps {
		if strings.Trim(app, "\n") != strings.Trim(program, "\n") {
			cleanedApps2 = append(cleanedApps2, app)
		} else {
			foundRunning = true
		}
	}
	if !foundRunning && !ignoredRunningFrom {
		statuser.ErrorMsg("\n"+strings.TrimSuffix(program, "\n")+" is not open", 1)
	}
	return cleanedApps2
}
