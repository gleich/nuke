package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CheckWithUser ... Make sure the user want to quit all application
func CheckWithUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you sure you want to force quite all\napplications? (y or n)")
	continueProgram, _ := reader.ReadString('\n')
	if strings.ToLower(continueProgram) == "n\n" {
		os.Exit(0)
	}
}
