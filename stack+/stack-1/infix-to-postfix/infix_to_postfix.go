package main

import (
	"fmt"
	"strings"
)

func solve(A string) string {
	var result strings.Builder
	var stack []rune

	// Precedence map: ( keeps 0 so operators don't pop it.
	precedence := map[rune]int{
		'^': 3,
		'*': 2, '/': 2,
		'+': 1, '-': 1,
		'(': 0,
	}

	for _, char := range A {
		switch {
		// 1. Operands: Add directly to result
		case char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z':
			result.WriteRune(char)

		// 2. Open Bracket: Push to stack
		case char == '(':
			stack = append(stack, char)

		// 3. Close Bracket: Pop until '(' is found
		case char == ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				result.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 { // Remove the '('
				stack = stack[:len(stack)-1]
			}

		// 4. Operators: Pop all >= precedence
		default:
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				// If top operator is >= current, it must be evaluated first
				if precedence[top] >= precedence[char] {
					result.WriteRune(top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, char)
		}
	}

	// 5. Final Drain: Empty the stack
	for len(stack) > 0 {
		result.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result.String()
}

func main() {
	// Example: "x^y/(a*z)+b" -> "xy^az*/b+"
	fmt.Println("Result 1:", solve("x^y/(a*z)+b"))

	// Example: "a+b*c" -> "abc*+"
	fmt.Println("Result 2:", solve("a+b*c"))
}
