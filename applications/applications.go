package applications

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Get ... Get a list of all the open application
func Get() []string {
	out, err := exec.Command("osascript", "-e", `tell application "System Events" to get name of (processes where background only is false)`).Output()
	if err != nil {
		log.Fatal(err)
	}
	dirtyApplications := strings.Split(string(out), ",")
	cleanApplications := []string{}
	for _, app := range dirtyApplications {
		if strings.Trim(strings.TrimSpace(app), "\n") != "Finder" {
			cleanApplications = append(cleanApplications, strings.TrimSpace(app))
		}
	}
	return cleanApplications
}

// Quit ... Quit an application
func Quit(name string) {
	cleanedName := strings.ReplaceAll(name, " ", "\\ ")
	cmd := exec.Command("pkill", "-x", cleanedName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("💣 " + name + " quitted")
}
