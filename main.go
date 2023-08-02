package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gleich/desktop"
	"github.com/gleich/nuke/pkg/config"
	"github.com/gleich/nuke/pkg/input"
	"github.com/gleich/nuke/pkg/macos"
	"github.com/gleich/nuke/pkg/output"
	"github.com/gleich/nuke/pkg/version"
	"github.com/gleich/statuser/v2"
)

func main() {
	statuser.Emojis = false
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

	if !configContents.IgnoreUpdates {
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
	cleanedApps := input.ExecutingTerm(apps, ignoredApps, configContents.IgnoredRunningFrom)
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
		statuser.Success("Quit " + app)
	}
	if operatingSystem == "darwin" {
		macos.CloseFinder()
	}

	fmt.Println()
	statuser.Success("Quit All Applications! Have a good day :)")
}
