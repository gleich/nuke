package main

import (
	"log"
	"runtime"

	"github.com/Matt-Gleich/nuke/applications"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/output"
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatal("Only support for macOS")
	}
	output.Title()
	apps := applications.GetApplications()
	input.ExecutingTerm(apps)
}
