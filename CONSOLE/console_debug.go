package CONSOLE

import (
	"fmt"

	CPU_6502 "github.com/cassianoperin/6502"
)

// Console debug command
func Console_Command_Debug(text_slice []string) {
	// Test the command syntax
	if len(text_slice) == 1 {

		// Show current value
		if CPU_6502.Debug {
			fmt.Printf("Debug status: Enabled\n\n")
		} else {
			fmt.Printf("Debug status: Disabled\n\n")
		}

	} else if len(text_slice) > 2 {

		// Print debug usage
		fmt.Printf("Usage: debug <on|off>\n\n")

	} else {

		if text_slice[1] == "on" || text_slice[1] == "off" {
			if text_slice[1] == "on" {
				CPU_6502.Debug = true
				fmt.Printf("Debug mode enabled\n\n")
			} else {
				CPU_6502.Debug = false
				fmt.Printf("Debug mode disabled\n\n")
			}
		} else {
			// Print debug usage
			fmt.Printf("Usage: debug <on|off>\n\n")
		}

	}
}
