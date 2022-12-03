package CONSOLE

import (
	"flag"
	"fmt"

	CPU_6502 "github.com/cassianoperin/6502_GO_Core"
)

// Console reset command
func Console_Command_Reset(text_slice []string) {
	// Test the command syntax
	if len(text_slice) > 1 {

		// Show current value
		fmt.Printf("Reset doesn't accept arguments\n\n")

	} else {

		// Reset CPU
		CPU_6502.Initialize()
		CPU_6502.Reset()

		// Load Program again into memory
		CPU_6502.ReadROM(flag.Arg(0))

		// Print the Header
		Console_PrintHeader()
	}
}
