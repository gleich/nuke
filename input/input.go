package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ExecutingTerm ... Ask the user what terminal they are executing from
func ExecutingTerm(runningApps []string) []string {
	for _, app := range runningApps {
		fmt.Println(app)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nOut of the applications above, which one\nis the one you are executing this command from?")
	program, _ := reader.ReadString('\n')
	var found bool
	cleanedApps := []string{}
	for _, app := range runningApps {
		if strings.Trim(app, "\n") == strings.Trim(program, "\n") {
			found = true
		} else {
			cleanedApps = append(cleanedApps, app)
		}
	}
	if !found {
		fmt.Println(program + " is not an open application")
		os.Exit(1)
	}
	return cleanedApps
}
