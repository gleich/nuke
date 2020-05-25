package applications

import (
	"log"
	"os/exec"
	"strings"
)

// GetApplications ... Get a list of all the open application
func GetApplications() []string {
	out, err := exec.Command(`osascript`, `-e`, `tell application "System Events" to get name of (processes where background only is false)`).Output()
	if err != nil {
		log.Fatal(err)
	}
	dirtyApplications := strings.Split(string(out), ",")
	cleanApplications := []string{}
	for _, app := range dirtyApplications {
		cleanApplications = append(cleanApplications, strings.Trim(app, " "))
	}
	return cleanApplications
}
