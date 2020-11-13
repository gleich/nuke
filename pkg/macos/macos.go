package macos

import (
	"os/exec"
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Close all open finder windows
func CloseFinder() {
	s := spinner.New(spinner.CharSets[13], 30*time.Millisecond)
	s.Suffix = " Closing Finder windows"
	s.Start()
	err := exec.Command("osascript", "-e", `tell application "Finder" to close windows`).Run()
	if err != nil {
		statuser.Error("Failed to close all finder windows", err, 1)
	}
	s.Stop()
	statuser.Success("Closed all Finder windows")
}
