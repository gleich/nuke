package main

import (
	"fmt"

	"github.com/Matt-Gleich/nuke/applications"
	"github.com/Matt-Gleich/nuke/input"
	"github.com/Matt-Gleich/nuke/output"
)

func main() {
	output.Title()
	input.CheckWithUser()
	fmt.Println(applications.GetApplications())
}
