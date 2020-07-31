package macos

import (
	"os/exec"

	"github.com/Matt-Gleich/nuke/output"
	"github.com/Matt-Gleich/statuser/v2"
)

// CloseFinder ... Close all open finder windows
func CloseFinder() {
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		statuser.Error("Failed to close all finder windows", err, 0)
	}
	output.Success("💥 Closed all Finder windows")
}
