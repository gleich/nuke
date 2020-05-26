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
	fmt.Println("Out of the applications above, which one\nis the one you are executing this command from?")
	program, _ := reader.ReadString('\n')
	var found bool
	cleanedApps := []string{}
	for _, app := range runningApps {
		if app == strings.TrimSuffix(program, "\n") {
			found = true
		}
		cleanedApps = append(cleanedApps, app)
	}
	if !found {
		fmt.Println(program + " is not an open application")
	}
	return cleanedApps
}
