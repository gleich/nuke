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
		fmt.Println("This application only supports macos")
	}
	output.Title()
	apps := applications.Get()
	cleanedApps := input.ExecutingTerm(apps)
	fmt.Println("")
	for _, app := range cleanedApps {
		applications.Quit(app)
	}
}
