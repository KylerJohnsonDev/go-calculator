package main

import (
	"bufio"
	"calculator/operations"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator is up and running \n Press 'enter' to continue...")
	fmt.Scanln()
	var operation_selection string

	for strings.Trim(operation_selection, "\r\n") != "exit" {
		fmt.Println("What would you like to do?\n1) Add\n2) Subtract\n3) Multiply\n4) Divide\nOr type \"exit\" to stop.")
		operation_selection, _ = reader.ReadString('\n')
		trimmed_op_selection := strings.Trim(operation_selection, "\r\n") // why \r\n?

		switch trimmed_op_selection {
		case "1":
			operations.Start_addition_operation(*reader)
		case "2":
			operations.Start_subtraction_operation(*reader)
		case "3":
			fmt.Println(fmt.Sprintf("You selected %v", trimmed_op_selection))
		case "4":
			fmt.Println(fmt.Sprintf("You selected %v", trimmed_op_selection))
		case "exit":
			fmt.Println(fmt.Sprintf("You selected %v", trimmed_op_selection))
		default:
			fmt.Println("Not a valid selection. Please select an operation from the menu.")
		}
	}
}
