package main

import (
	"fmt"
)

func Sqrt(A int) int {
	// Base case for 0 and 1
	if A < 2 {
		return A
	}

	l, r := 0, A
	ans := 0

	for l <= r {
		mid := l + (r-l)/2

		// Use int64 to prevent overflow when calculating mid * mid
		midSq := int64(mid) * int64(mid)

		if midSq == int64(A) {
			return mid
		}

		if midSq < int64(A) {
			// mid is a potential floor answer
			ans = mid
			l = mid + 1
		} else {
			// mid is too big
			r = mid - 1
		}
	}

	return ans
}

func main() {
	fmt.Println("Sqrt(11):", Sqrt(11)) // Output: 3
	fmt.Println("Sqrt(9):", Sqrt(9))   // Output: 3
	fmt.Println("Sqrt(0):", Sqrt(0))   // Output: 0
	fmt.Println("Sqrt(1):", Sqrt(1))   // Output: 1
}
