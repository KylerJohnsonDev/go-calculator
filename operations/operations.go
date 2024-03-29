package operations

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parse_number_inputs(number_input string) []int64 {
	trimmed_input := strings.Trim(number_input, "\r\n")
	pre_parsed_nums := strings.Split(trimmed_input, ",")
	var nums []int64
	for _, num_string := range pre_parsed_nums {
		trimmed := strings.Trim(num_string, " ")
		number, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to parse string input into numbers. Input: %v", number_input))
			fmt.Println(err)
		}
		nums = append(nums, number)
	}
	return nums
}

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
	nums_to_add := parse_number_inputs(nums_to_add_input)
	sum := add_numbers(nums_to_add)
	fmt.Println(fmt.Sprintf("Sum: %v", sum))
	fmt.Scanln()
}

func subtract_numbers(numbers []int64) int64 {
	var difference int64 = 0
	for index, num := range numbers {
		if index == 0 {
			difference = num
		} else {
			difference = difference - num
		}
	}
	return difference
}

func Start_subtraction_operation(reader bufio.Reader) {
	fmt.Println("Please type all of the numbers you wish to subtract in order, separated by a comma below.")
	nums_to_subtract_input, _ := reader.ReadString('\n')
	nums_to_subtract := parse_number_inputs(nums_to_subtract_input)
	difference := subtract_numbers(nums_to_subtract)
	fmt.Println(fmt.Sprintf("Difference: %v", difference))
	fmt.Scanln()
}

func multiply_numbers(numbers []int64) int64 {
	var product int64 = 0
	for index, num := range numbers {
		if index == 0 {
			product = num
		} else {
			product = product * num
		}
	}
	return product
}

func Start_multiplication_operation(reader bufio.Reader) {
	fmt.Println("Please type all of the numbers you wish to multiply, separated by a comma below.")
	nums_to_multiply_input, _ := reader.ReadString('\n')
	nums_to_multiply := parse_number_inputs(nums_to_multiply_input)
	product := multiply_numbers(nums_to_multiply)
	fmt.Println(fmt.Sprintf("Product: %v", product))
	fmt.Scanln()
}

func divide_numbers(numbers []int64) int64 {
	var quotient int64 = 0
	for index, num := range numbers {
		if index == 0 {
			quotient = num
		} else if num == 0 {
			fmt.Println("Error: cannot divide by 0")
			return 0
		} else {
			quotient = quotient / num
		}
	}
	return quotient
}

func Start_division_operation(reader bufio.Reader) {
	fmt.Println("Please type all of the numbers you wish to divide in order, separated by a comma below.")
	nums_to_divide_input, _ := reader.ReadString('\n')
	nums_to_divide := parse_number_inputs(nums_to_divide_input)
	quotient := divide_numbers(nums_to_divide)
	fmt.Println(fmt.Sprintf("Quotient: %v", quotient))
	fmt.Scanln()
}
