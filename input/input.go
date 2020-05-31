package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Matt-Gleich/nuke/output"
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
		output.Error("\n" + strings.TrimSuffix(program, "\n") + " is not open")
	}
	return cleanedApps
}
