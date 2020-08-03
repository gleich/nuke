package version

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Matt-Gleich/statuser/v2"
)

// CheckForUpdate ... Check for an update for the CLI
func CheckForUpdate() {
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
	json.NewDecoder(resp.Body).Decode(&result)

	if result["tag_name"] != "v4.1.1" {
		fmt.Println("THERE IS AN UPDATE AVALIABLE")
		fmt.Println("PLEASE UPDATE ASAP\n")
	}
}
