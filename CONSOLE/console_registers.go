package CONSOLE

import (
	"fmt"
	"strings"

	CPU_6502 "github.com/cassianoperin/6502_GO_Core"
)

// Console registers command
func Console_Command_Registers(text_slice []string) {

	var location_value []string

	// Test the command syntax
	if len(text_slice) == 1 || len(text_slice) > 2 {

		// Print usage
		fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")
	} else {

		// After, split the argument in LOCATION=VALUE
		location_value = strings.Split(text_slice[1], "=")

		if len(location_value) == 1 || len(location_value) > 2 || location_value[1] == "" || location_value[0] == "" {

			// Print usage
			fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")

		} else {

			location := strings.ToUpper(location_value[0])

			// Validate the value of locations
			if location == "PC" || location == "A" || location == "X" || location == "Y" {

				// Check if input is Decimar of Hexadecimal and convert to integer
				mem_arg, error_flag := Console_Hex_or_Dec(location_value[1])

				if !error_flag {
					// Value limits
					if location == "PC" {
						if mem_arg <= 65535 && mem_arg >= 0 {
							CPU_6502.PC = uint16(mem_arg)
							fmt.Printf("\tPC set to 0x%02X (%d)\n\n", CPU_6502.PC, CPU_6502.PC)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFFFF (65536)\n\n")
						}
					}

					if location == "A" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CPU_6502.A = byte(mem_arg)
							fmt.Printf("\tA set to 0x%02X (%d)\n\n", CPU_6502.A, CPU_6502.A)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

					if location == "X" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CPU_6502.X = byte(mem_arg)
							fmt.Printf("\tX set to 0x%02X (%d)\n\n", CPU_6502.X, CPU_6502.X)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

					if location == "Y" {
						if mem_arg <= 255 && mem_arg >= 0 {
							CPU_6502.Y = byte(mem_arg)
							fmt.Printf("\tY set to %02X (%d)\n\n", CPU_6502.Y, CPU_6502.Y)
							Console_PrintHeader()
						} else {
							fmt.Printf("Invalid Address. Should be in range 0x0000 and 0xFF (255)\n\n")
						}
					}

				} else {
					fmt.Printf("Invalid value %s\n\n", location_value[1])
				}

			} else {

				// Print usage
				fmt.Printf("Usage: registers <A|X|Y|PC>=<value>\n\n")
			}

		}

	}

}
