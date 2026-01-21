package main

import "fmt"

func getPSE(A []int) []int {
	n := len(A)

	stack := make([]int, 0) // store index, not element
	pse := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && A[stack[len(stack)-1]] >= A[i] {
			stack = stack[:len(stack)-1]
		}

		pse[i] = -1
		if len(stack) > 0 {
			pse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return pse
}

func getNSEE(A []int) []int {
	n := len(A)

	stack := make([]int, 0) // store index, not element
	nse := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack)-1]] > A[i] {
			stack = stack[:len(stack)-1]
		}

		nse[i] = n
		if len(stack) > 0 {
			nse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return nse
}

func getPGE(A []int) []int {
	n := len(A)

	stack := make([]int, 0) // store index, not element
	pge := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && A[stack[len(stack)-1]] <= A[i] {
			stack = stack[:len(stack)-1]
		}

		pge[i] = -1
		if len(stack) > 0 {
			pge[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return pge
}

func getNGEE(A []int) []int {
	n := len(A)

	stack := make([]int, 0) // store index, not element
	nge := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack)-1]] < A[i] {
			stack = stack[:len(stack)-1]
		}

		nge[i] = n
		if len(stack) > 0 {
			nge[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return nge
}

/*
To handle duplicates

	Strict vs. Non-Strict Inequality
		To ensure every subarray is counted exactly once, we use a "Tie-Breaking" rule.
		We treat the element that appears earlier as the "dominant" one.
			- On the Left side, we look for the nearest element that is strictly smaller/greater.
			- On the Right side, we look for the nearest element that is smaller/greater or equal.
			Note: It doesn't matter which side is strict, as long as one is strict and the other is not.
*/
func solve(A []int) int {
	const MOD int64 = 1e9 + 7

	// an element is MINIMUM in subarrays upto the next & previous smaller element on both sides
	nse := getNSEE(A)
	pse := getPSE(A)

	// an element is MAXIMUM in subarrays upto the next & previous greater element on both sides
	nge := getNGEE(A)
	pge := getPGE(A)

	var total int64 = 0
	for i := 0; i < len(A); i++ {
		minSub := int64(i-pse[i]) * int64(nse[i]-i)
		maxSub := int64(i-pge[i]) * int64(nge[i]-i)

		minContribution := (minSub % MOD * int64(A[i]) % MOD) % MOD
		maxContribution := (maxSub % MOD * int64(A[i]) % MOD) % MOD

		total = (total + maxContribution - minContribution + int64(MOD)) % MOD
	}

	return int(total)
}

func main() {
	// Case 1: [4, 7, 3, 8]
	A := []int{4, 7, 3, 8}
	fmt.Printf("Input: %v | Output: %v\n", A, solve(A))
	// Expected: 26
}
