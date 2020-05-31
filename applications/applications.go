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
	cmd := exec.Command("pkill", "-x", cleanedName)
	err := cmd.Run()
	if err != nil {
		output.Error("Failed to quit " + name)
	}
	output.Success("💥 Quitted " + name)
}
