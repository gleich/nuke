package output

import (
	"github.com/wayneashleyberry/truecolor/pkg/color"
)

// Title ... Output the application title
func Title() {
	title := []string{
		"███╗   ██╗ ██╗   ██╗ ██╗  ██╗ ███████╗",
		"████╗  ██║ ██║   ██║ ██║ ██╔╝ ██╔════╝",
		"██╔██╗ ██║ ██║   ██║ █████╔╝  █████╗",
		"██║╚██╗██║ ██║   ██║ ██╔═██╗  ██╔══╝",
		"██║ ╚████║ ╚██████╔╝ ██║  ██╗ ███████╗",
		"╚═╝  ╚═══╝  ╚═════╝  ╚═╝  ╚═╝ ╚══════╝",
	}
	for i, value := range title {
		color.Color(216, uint8(i*30), 0).Println(value)
	}
}
