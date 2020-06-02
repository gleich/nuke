package main

import (
	"fmt"
	"runtime"

	"github.com/Matt-Gleich/nuke/applications"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/output"
)

func main() {
	if runtime.GOOS != "darwin" {
		output.Error("This application only runs on macos. Your running on " + runtime.GOOS)
	}
	output.Title()
	apps := applications.Get()
	cleanedApps := input.ExecutingTerm(apps)
	fmt.Println("")
	for _, app := range cleanedApps {
		applications.Quit(app)
	}
	output.Success("\n🤯 All Applications Quitted!")
}
