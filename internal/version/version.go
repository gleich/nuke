package version

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Check for an update for the CLI
func CheckForUpdate() {
	s := spinner.New(spinner.CharSets[13], 30*time.Millisecond)
	s.Suffix = " Checking for update"
	s.Start()

	// Checking network connection
	checkResp, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return
	}
	defer checkResp.Body.Close()

	// Making the actual request
	resp, err := http.Get("https://api.github.com/repos/Matt-Gleich/nuke/releases/latest")
	if err != nil {
		statuser.Error("Failed to get latest version from GitHub", err, 1)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		statuser.Error("Failed to get latest version from GitHub", err, 1)
	}

	if result["tag_name"] != "v4.2.2" {
		fmt.Println("THERE IS AN UPDATE AVALIABLE")
		fmt.Println(`PLEASE UPDATE ASAP
		`)
	}
	s.Stop()
}
