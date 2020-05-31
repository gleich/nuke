package output

import (
	"os"

	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Error ... Output an error
func Error(message string) {
	color.Color(255, 0, 0).Println(message)
	os.Exit(0)
}

// Success ... Output a success
func Success(message string) {
	color.Color(0, 255, 0).Println(message)
}
