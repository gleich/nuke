package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// ExecutingTerm ... Ask the user what terminal they are executing from
func ExecutingTerm(runningApps, ignoredApps []string) []string {
	for _, app := range runningApps {
		color.Color(0, 255, 0).Println(app)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nOut of the open applications above, which one\nis the one you are executing this command from?")
	program, _ := reader.ReadString('\n')
	var found bool
	cleanedApps := []string{}
	for _, app := range runningApps {
		if strings.Trim(app, "\n") == strings.Trim(program, "\n") {
			found = true
		} else {
			var notIgnored bool
			for _, ignoredApp := range ignoredApps {
				notIgnored = ignoredApp != app
				if notIgnored {
					break
				}
			}
			if notIgnored || len(ignoredApps) == 0 {
				cleanedApps = append(cleanedApps, app)
			}
		}
	}
	if !found {
		statuser.ErrorMsg("\n"+strings.TrimSuffix(program, "\n")+" is not open", 0)
	}
	return cleanedApps
}
