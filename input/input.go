package input

import (
	"sort"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/manifoldco/promptui"
)

// ExecutingTerm ... Ask the user what terminal they are executing from
func ExecutingTerm(runningApps, ignoredApps []string) []string {
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
	prompt := promptui.Select{
		Label: "Executing nuke command from",
		Items: cleanedApps,
	}
	_, program, err := prompt.Run()
	if err != nil {
		statuser.Error("Failed to get executing terminal", err, 1)
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
	if !foundRunning {
		statuser.ErrorMsg("\n"+strings.TrimSuffix(program, "\n")+" is not open", 1)
	}
	return cleanedApps2
}
