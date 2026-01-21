package main

import (
	"fmt"
)

// MaxSlidingWindow returns the maximum element in each sliding window of size k.
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Result will have size n - k + 1
	result := make([]int, 0, n-k+1)

	// deque will store INDICES of the elements
	// We maintain a monotonic decreasing order in the deque
	var deque []int

	for i := 0; i < n; i++ {
		// 1. Remove indices that are out of the current window range [i-k+1, i]
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:] // Pop from the front
		}

		// 2. Maintain monotonic decreasing order:
		// Remove elements from the back of the deque that are smaller than the current element.
		// These elements can never be the maximum because nums[i] is larger and stays in the window longer.
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1] // Pop from the back
		}

		// 3. Add the current index to the back
		deque = append(deque, i)

		// 4. If the window has reached size k, the result storing process has started - the front of the deque is our maximum
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("Sliding Window Maximums:", MaxSlidingWindow(nums, k))
	// Expected Output: [3, 3, 5, 5, 6, 7]
}
