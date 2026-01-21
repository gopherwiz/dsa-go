package main

import (
	"fmt"
	"sort"
)

func FindPlatform(arr []int, dep []int) int {
	// 1. Sort both independently
	sort.Ints(arr)
	sort.Ints(dep)

	n := len(arr)
	platformsNeeded := 0
	maxPlatforms := 0

	i := 0 // Arrival pointer
	j := 0 // Departure pointer

	// 2. Sweep through the timeline
	for i < n && j < n {
		if arr[i] <= dep[j] {
			// A train is arriving before the next one leaves
			platformsNeeded++
			i++
		} else {
			// A train leaves, freeing a platform
			platformsNeeded--
			j++
		}

		// 3. Update the peak occupancy
		if platformsNeeded > maxPlatforms {
			maxPlatforms = platformsNeeded
		}
	}

	return maxPlatforms
}

func main() {
	// Input: Times represented in 24H format (e.g., 900 = 9:00 AM)
	arrivals := []int{900, 940, 950, 1100, 1500, 1800}
	departures := []int{910, 1200, 1120, 1130, 1900, 2000}

	result := FindPlatform(arrivals, departures)

	fmt.Printf("Arrival Times:   %v\n", arrivals)
	fmt.Printf("Departure Times: %v\n", departures)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Minimum Platforms Required: %d\n", result)
}
