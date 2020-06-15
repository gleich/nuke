package main

import (
	"fmt"
	"runtime"

	"github.com/Matt-Gleich/nuke/applications"
	"github.com/Matt-Gleich/nuke/config"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/output"
)

func main() {
	if runtime.GOOS != "darwin" {
		output.Error("This application only runs on macos. Your running on " + runtime.GOOS)
	}
	var ignoredApps []string
	if config.Exists() {
		ignoredApps = config.Read()["ignored"]
	}
	output.Title()
	apps := applications.Get()
	cleanedApps := input.ExecutingTerm(apps, ignoredApps)
	fmt.Println("")
	for _, app := range cleanedApps {
		applications.Quit(app)
	}
	applications.CloseFinder()
	output.Success("\n🤯 All Applications Quitted!")
}
