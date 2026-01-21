package main

import (
	"fmt"
	"strconv"
)

// evalRPN evaluates the value of an arithmetic expression in Reverse Polish Notation.
func evalRPN(A []string) int {
	// A slice used as a stack to store operands.
	var stack []int

	for _, token := range A {
		switch token {
		case "+", "-", "*", "/":
			// Operators: Pop the top two elements.
			// The first pop is the second operand; the second pop is the first operand.
			if len(stack) < 2 {
				continue // Should not happen in valid RPN
			}

			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // "Pop" both elements

			var result int
			switch token {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2 // Go integer division naturally truncates toward zero
			}
			// Push the result back onto the stack
			stack = append(stack, result)

		default:
			// Operand: Convert string to integer and push onto stack.
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	// The final result is the only element remaining in the stack.
	return stack[0]
}

func main() {
	// Example 1: ["2", "1", "+", "3", "*"] -> (2 + 1) * 3 = 9
	input1 := []string{"2", "1", "+", "3", "*"}
	fmt.Printf("Input: %v | Result: %d\n", input1, evalRPN(input1))

	// Example 2: ["4", "13", "5", "/", "+"] -> 4 + (13 / 5) = 4 + 2 = 6
	input2 := []string{"4", "13", "5", "/", "+"}
	fmt.Printf("Input: %v | Result: %d\n", input2, evalRPN(input2))
}
