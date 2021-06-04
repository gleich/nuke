package input

import (
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/statuser/v2"
)

// Ask the user what terminal they are executing from
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

	program := struct {
		Program string `survey:"program"`
	}{}
	err := survey.Ask([]*survey.Question{{
		Name: "program",
		Prompt: &survey.Select{
			Message:  "Running nuke from (what app your currently using)",
			Options:  cleanedApps,
			PageSize: 25,
		},
	}}, &program)
	if err != nil {
		statuser.Error("Failed to get running terminal", err, 1)
	}

	var foundRunning bool
	cleanedApps2 := []string{}
	for _, app := range cleanedApps {
		if strings.Trim(app, "\n") != strings.Trim(program.Program, "\n") {
			cleanedApps2 = append(cleanedApps2, app)
		} else {
			foundRunning = true
		}
	}
	if !foundRunning {
		statuser.ErrorMsg("\n"+strings.TrimSuffix(program.Program, "\n")+" is not open", 1)
	}
	return cleanedApps2
}
