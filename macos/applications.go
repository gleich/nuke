package macos

import (
	"os/exec"
	"strings"

	"github.com/Matt-Gleich/nuke/output"
	"github.com/Matt-Gleich/statuser/v2"
)

// GetApplications ... Get a list of all the open application
func GetApplications() []string {
	out, err := exec.Command("osascript", "-e", `tell application "System Events" to get name of (processes where background only is false)`).Output()
	if err != nil {
		err := exec.Command("osascript", "-e", `tell application "System Events" to activate`).Run()
		if err != nil {
			statuser.Error("Failed to get running list of applications", err, 0)
		}
	}
	dirtyApplications := strings.Split(string(out), ",")
	cleanApplications := []string{}
	for _, app := range dirtyApplications {
		if strings.TrimSuffix(strings.TrimSpace(app), "\n") != "Finder" {
			cleanApplications = append(cleanApplications, strings.TrimSpace(app))
		}
	}
	return cleanApplications
}

// QuitApp ... Quit an application
func QuitApp(name string) {
	cleanedName := strings.ReplaceAll(name, " ", "\\ ")
	err := exec.Command("pkill", "-x", cleanedName).Run()
	if err != nil {
		statuser.Error("Failed to quit "+name, err, 0)
	}
	output.Success("💥 Quitted " + name)
}

// CloseFinder ... Close all open finder windows
func CloseFinder() {
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		statuser.Error("Failed to close all finder windows", err, 0)
	}
	output.Success("💥 Closed all Finder windows")
}
