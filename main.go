package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/Matt-Gleich/desktop"
	"github.com/Matt-Gleich/nuke/internal/config"
	"github.com/Matt-Gleich/nuke/internal/input"
	"github.com/Matt-Gleich/nuke/internal/macos"
	"github.com/Matt-Gleich/nuke/internal/output"
	"github.com/Matt-Gleich/nuke/internal/version"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

func main() {
	operatingSystem := runtime.GOOS
	if operatingSystem != "darwin" && operatingSystem != "linux" {
		statuser.ErrorMsg("This application only works on macOS and linux", 1)
	}

	var configContents config.Conf
	if config.Exists() {
		config.Read(&configContents)
	}

	// Ignoring apps
	var ignoredApps []string
	if config.Exists() {
		ignoredApps = configContents.IgnoredApps
	}
	ignoredApps = append(ignoredApps, os.Args[1:]...)

	output.Title()

	if operatingSystem == "linux" && !configContents.IgnoreUpdates {
		version.CheckForUpdate()
	}

	// Getting running applications
	var apps []string
	var appsWithPIDs map[string]int
	s := spinner.New(spinner.CharSets[13], 30*time.Millisecond)
	s.Suffix = " Getting list of applications"
	s.Start()
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
		for app := range linuxApps {
			apps = append(apps, app)
		}
		appsWithPIDs = linuxApps
	}
	s.Stop()

	// Getting executing terminal
	cleanedApps := input.ExecutingTerm(apps, ignoredApps)
	fmt.Println("")

	// Quitting applications
	for _, app := range cleanedApps {
		s2 := spinner.New(spinner.CharSets[13], 30*time.Millisecond)
		s2.Suffix = " Quitting " + app
		s2.Start()
		switch operatingSystem {
		case "darwin":
			err := desktop.MacOSQuitApp(app)
			if err != nil {
				statuser.Error("Failed to quit "+app, err, 1)
			}
		case "linux":
			err := desktop.LinuxQuitApp(appsWithPIDs[app])
			if err != nil {
				statuser.Error("Failed to quit "+app, err, 1)
			}
		}
		s2.Stop()
		output.Success("💥 Quitted " + app)
	}
	if operatingSystem == "darwin" {
		macos.CloseFinder()
	}
	output.Success("\n🤯 All Applications Quitted!")
}
