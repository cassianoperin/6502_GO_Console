package main

import (
	"flag"
	"fmt"

	"github.com/cassianoperin/6502_console/CLI"
	"github.com/cassianoperin/6502_console/CONSOLE"

	CPU_6502 "github.com/cassianoperin/6502"
)

func main() {

	fmt.Printf("\nMOS 6502 CPU Emulator\n\n")

	// Validate the Arguments
	CLI.CheckArgs()

	// Set initial variables values
	CPU_6502.Initialize()

	// Initialize Timers
	CPU_6502.InitializeTimers()

	// Read ROM to the memory
	CPU_6502.ReadROM(flag.Arg(0))
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_functional_test.bin")
	// readROM("/Users/cassiano/go/src/6502/TestPrograms/6502_decimal_test.bin")

	// Reset system
	CPU_6502.Reset()

	// Start Console Mode
	CONSOLE.StartConsole()

}
