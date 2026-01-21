package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start int
	end   int
}

func MaxActivities(start []int, end []int) int {
	n := len(start)
	if n == 0 {
		return 0
	}

	// 1. Combine and Sort by End Time
	activities := make([]Activity, n)
	for i := 0; i < n; i++ {
		activities[i] = Activity{start[i], end[i]}
	}

	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	})

	// 2. Select activities
	count := 1 // Always pick the first one (finishes earliest)
	lastFinishTime := activities[0].end

	for i := 1; i < n; i++ {
		// If this activity starts after the last one finished
		if activities[i].start >= lastFinishTime {
			count++
			lastFinishTime = activities[i].end
		}
	}

	return count
}

func main() {
	start := []int{1, 3, 0, 5, 8, 5}
	end := []int{2, 4, 6, 7, 9, 9}

	result := MaxActivities(start, end)
	fmt.Printf("Maximum number of non-conflicting activities: %d\n", result)
}
