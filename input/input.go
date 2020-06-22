package input

import (
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/manifoldco/promptui"
)

// ExecutingTerm ... Ask the user what terminal they are executing from
func ExecutingTerm(runningApps, ignoredApps []string) []string {
	prompt := promptui.Select{
		Label: "Executing nuke command from",
		Items: runningApps,
	}
	_, program, err := prompt.Run()
	if err != nil {
		statuser.Error("Failed to get executing terminal", err, 1)
	}
	var found bool
	cleanedApps := []string{}
	for _, app := range runningApps {
		if strings.Trim(app, "\n") == strings.Trim(program, "\n") {
			found = true
		} else {
			var notIgnored bool
			for _, ignoredApp := range ignoredApps {
				notIgnored = ignoredApp != app
				if notIgnored {
					break
				}
			}
			if notIgnored || len(ignoredApps) == 0 {
				cleanedApps = append(cleanedApps, app)
			}
		}
	}
	if !found {
		statuser.ErrorMsg("\n"+strings.TrimSuffix(program, "\n")+" is not open", 1)
	}
	return cleanedApps
}
