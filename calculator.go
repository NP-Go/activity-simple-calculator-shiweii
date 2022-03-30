package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int {
	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	return a / b
}

// Read user input and covert string to integer, question will be repeated if non integer is entered.
func intInputCheck(i string) int {
	var ret int
	for {
		fmt.Println("Enter an integer: ")
		fmt.Scanln(&i)
		value, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Please enter integer value only.")
		} else {
			ret = value
			break
		}
	}
	return ret
}

// Helper function to check if user enter valid arithmetic process.
func contains(a []string, str string) bool {
	for _, v := range a {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	var a, b, result int
	var input1, input2, process string

	a = intInputCheck(input1)

	// Check if valid arithmetic process is entered.
	for {
		fmt.Println("Enter process: (add, sub, mul, div)")
		fmt.Scanln(&process)
		// Added + - x / for usuability
		arithmetic := []string{"add", "sub", "mul", "div", "+", "-", "x", "X", "*", "/"}
		if !contains(arithmetic, process) {
			fmt.Println("Please enter process: (add, sub, mul, div) only")
		} else {
			break
		}
	}

	b = intInputCheck(input2)

	//Insert code here

	// Switch case to determine which arithmetic function to process
	switch process {
	case "+", "add":
		result = add(a, b)
		fmt.Println("Result: ", result)
	case "-", "sub":
		result = subtract(a, b)
		fmt.Println("Result: ", result)
	case "x", "X", "mul", "*":
		result = multiply(a, b)
		fmt.Println("Result: ", result)
	case "/", "div":
		if b == 0 {
			fmt.Println("Unable to divide by zero.")
		} else {
			result = divide(a, b)
			fmt.Println("Result: ", result)
		}
	}
}
