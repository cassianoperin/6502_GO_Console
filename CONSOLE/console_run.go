package CONSOLE

import (
	"fmt"

	CPU_6502 "github.com/cassianoperin/6502_GO_Core"
)

// Console run command
func Console_Command_Run(text_slice []string) {
	var (
		current_PC      uint16
		step_count      int    = 0
		loop_count      uint16 = 0
		breakpoint_flag bool
	)

	// Test the command syntax
	if len(text_slice) > 1 {

		// Print run_limit usage
		fmt.Printf("run command doesn't accept arguments.\n\n")

	} else {

		for loop_count < loop_detection {

			// -------------------------- Start Checks --------------------------- //

			// Check Run step limits
			if step_count > run_limit {
				break // Exit for loop
			}

			// Check Breakpoints
			breakpoint_flag = Console_Check_breakpoints(breakpoint_flag)

			// Exit for loop if breakpoint has been found
			if breakpoint_flag {
				break
			}

			// -------------- Finish checks and return to execution -------------- //
			current_PC = CPU_6502.PC

			select {
			case <-second_timer: // Show the header and debug each second

				// Execute one instruction
				Console_Step(opcode_map, text_slice[0])

			default: // Just run the CPU

				// Execute one instruction without print
				Console_Step_without_debug(opcode_map, text_slice[0])

			}

			// Increase steps count
			step_count++

			// Check for run_limit and print debug message prior to quit loop
			if step_count > run_limit { // Print limit reached message
				fmt.Printf("Run limit reached (%d)\n\n", run_limit)
			}

			// Increase the loop counter
			if current_PC == CPU_6502.PC {
				loop_count++
			}

			// Check for loops and print debug message prior to quit loop
			if loop_count >= loop_detection {
				fmt.Printf("Loop detected on PC=%04X (%d repetitions)\n", CPU_6502.PC, loop_detection)
			}

		}

		// Print Header
		Console_PrintHeader()
	}

}
