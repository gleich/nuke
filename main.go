package main

import (
	"fmt"

	"github.com/Matt-Gleich/nuke/applications"
	"github.com/Matt-Gleich/nuke/input"
)

func main() {
	input.CheckWithUser()
	fmt.Println(applications.GetApplications())
}
