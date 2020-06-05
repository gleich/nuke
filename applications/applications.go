package applications

import (
	"log"
	"os/exec"
	"strings"

	"github.com/Matt-Gleich/nuke/output"
)

// Get ... Get a list of all the open application
func Get() []string {
	out, err := exec.Command("osascript", "-e", `tell application "System Events" to get name of (processes where background only is false)`).Output()
	if err != nil {
		log.Fatal(err)
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

// Quit ... Quit an application
func Quit(name string) {
	cleanedName := strings.ReplaceAll(name, " ", "\\ ")
	err := exec.Command("pkill", "-x", cleanedName).Run()
	if err != nil {
		output.Error("Failed to quit " + name)
	}
	output.Success("💥 Quitted " + name)
}

// CloseFinder ... Close all open finder windows
func CloseFinder() {
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		output.Error("Failed to close all finder windows")
	}
	output.Success("💥 Closed all Finder windows")
}
