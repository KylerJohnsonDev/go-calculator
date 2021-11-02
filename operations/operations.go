package operations

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func add_numbers(numbers []int64) int64 {
	var sum int64 = 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}

func Start_addition_operation(reader bufio.Reader) {
	fmt.Println("Please type all of the numbers you wish to add, separated by a comma below.")
	nums_to_add_input, _ := reader.ReadString('\n')
	trimmed_input := strings.Trim(nums_to_add_input, "\r\n")
	nums_to_add_text := strings.Split(trimmed_input, ",")
	var nums_to_add []int64
	for _, num_string:= range nums_to_add_text {
		trimmed := strings.Trim(num_string, " ")
		number, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to parse string input into numbers. Input: %v", nums_to_add_input))
			fmt.Println(err)
			return
		}
		nums_to_add = append(nums_to_add, number)
	}
	sum := add_numbers(nums_to_add)
	fmt.Println(fmt.Sprintf("Sum: %v", sum))
	fmt.Scanln()
}