package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Matt-Gleich/nuke/config"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/macos"
	"github.com/Matt-Gleich/nuke/output"
	"github.com/Matt-Gleich/statuser/v2"
)

func main() {
	operatingSystem := runtime.GOOS

	if operatingSystem != "darwin" || operatingSystem != "linux" {
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
		apps = macos.GetApplications()
	case "linux":
		fmt.Println("a")
	}

	// Getting executing terminal
	cleanedApps := input.ExecutingTerm(apps, ignoredApps)
	fmt.Println("")

	// Quitting applications
	for _, app := range cleanedApps {
		switch operatingSystem {
		case "darwin":
			macos.QuitApp(app)
		case "linux":
			fmt.Println(app)
		}
	}
	if operatingSystem == "darwin" {
		macos.CloseFinder()
	}
	output.Success("\n🤯 All Applications Quitted!")
}
