package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Matt-Gleich/desktop"
	"github.com/Matt-Gleich/nuke/config"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/macos"
	"github.com/Matt-Gleich/nuke/output"
	"github.com/Matt-Gleich/statuser/v2"
)

func main() {
	operatingSystem := runtime.GOOS
	if operatingSystem != "darwin" && operatingSystem != "linux" {
		statuser.ErrorMsg("This application only works on macOS and linux", 1)
	}

	// Ignoring apps
	var ignoredApps []string
	if config.Exists() {
		ignoredApps = config.Read()["ignored"]
	}
	for _, app := range os.Args[1:] {
		ignoredApps = append(ignoredApps, app)
	}

	output.Title()

	// Getting running applications
	var apps []string
	switch operatingSystem {
	case "darwin":
		macApps, err := desktop.MacOSApplications()
		if err != nil {
			statuser.Error("Failed to get macos applications", err, 1)
		}
		apps = macApps
	case "linux":
		linuxApps, err := desktop.LinuxApplications()
		if err != nil {
			statuser.Error("Failed to get linux applications", err, 1)
		}
		apps = linuxApps
	}

	// Getting executing terminal
	cleanedApps := input.ExecutingTerm(apps, ignoredApps)
	fmt.Println("")

	// Quitting applications
	for _, app := range cleanedApps {
		switch operatingSystem {
		case "darwin":
			desktop.MacOSQuitApp(app)
		case "linux":
			desktop.LinuxQuitApp(app)
		}
	}
	if operatingSystem == "darwin" {
		macos.CloseFinder()
	}
	output.Success("\n🤯 All Applications Quitted!")
}
