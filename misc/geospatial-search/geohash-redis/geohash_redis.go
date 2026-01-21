/*
REDIS GEOSPATIAL & SKIP LIST STUDY NOTES (2026)
==============================================

1. REDIS GEOSPATIAL INTERNALS
   - Data Structure: Redis uses a Sorted Set (ZSET) to store geospatial data.
   - Geohash Conversion: Coordinates (lat/long) are converted into a 52-bit integer score
     using a Z-order curve (space-filling curve).
   - Storage Logic: This integer is stored as the "score." This flattens 2D space into
     a 1D line, making search operations possible via range queries.
   - Complexity: GEOADD: O(9*log N); GEOSEARCH: O(9*log N + M);
   - 9*log N is the cost of using the Skip-List to jump to the relevant spatial grid
   - M is the cost of scanning through the M nearby candidates found within that spatial grid

2. THE SKIP LIST (THE "ENGINE")
   - Definition: A probabilistic data structure with a base linked list and multiple
     "express lanes" (layers) on top.
   - The "Tower": Each node contains an array of forward pointers. The size of this
     array is its "height."
   - Slipping/Skipping Logic: Start at the top; move Right if the next node is smaller
     than target; Drop Down if it is larger.

3. SKIP LIST VS. BALANCED TREES (RED-BLACK/AVL)
   - Range Queries: Superior in Skip Lists. Find the start, then perform a linear scan.
   - Insert/Delete: O(log N) without complex "rotations" or rigid rebalancing rules.
   - Memory: More flexible; averages ~1.33 pointers per node vs. fixed 2.0 in trees.
   - Logic: Uses probabilistic "coin flips" instead of deterministic rigid balancing.

4. WHY NOT AN ARRAYLIST?
   - Insertion Cost: Arrays require O(N) shifting of elements; Skip Lists only require
     updating O(1) pointers once the position is found.
   - Resizing: Arrays cause latency spikes during reallocation; Skip Lists grow
     node-by-node incrementally.

5. KEY TECHNICAL TERMS
   - Z-Order Curve: Math used to map 2D coordinates into a 1D integer while preserving locality.
   - Cache Locality: Arrays are better for reads, but Skip Lists are better for the
     dynamic updates required by a high-speed database.
   - ZipList: A compact array structure used by Redis for *very small* sets before
     converting to a Skip List.

6. GOLANG IMPLEMENTATION CONCEPT
   - Represented by the code below.
   - The "Tower" is the `Next []*Node` slice.
   - The height is `len(Next)`.
*/

/*
	REDIS GEO (GEOHASH / Z-ORDER CURVE)
	-----------------------------------
	- Concept: A hybrid "Middle Ground" approach. It uses bit-interleaving
	to create a Z-Order curve index.
	- Storage Location: In-Memory (RAM), but designed for easy persistence
	via snapshots (RDB) or Append-Only Files (AOF).
	- Data Structure: Sorted Set (ZSET). Internally, Redis uses a Skip-List
	where the "Score" is the 52-bit Geohash integer.
	- Disk Behavior: Linear Persistence. Since the data is a "Sorted Set,"
	Redis flattens the skip-list into a linear file during snapshots,
	making it faster to save/load than a pointer-heavy tree.
	- Recursion Point: Happens during Bitwise Interleaving. While the
	quadrant-within-quadrant logic is conceptually recursive, it is
	implemented using fast bit-shifts (<<, |, &) instead of function calls.
*/

package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel = 16
	p        = 0.5
)

type Node struct {
	Value int
	Next  []*Node // The "Tower"
}

type SkipList struct {
	Head *Node
}

func NewSkipList() *SkipList {
	return &SkipList{Head: &Node{Next: make([]*Node, maxLevel)}}
}

func (sl *SkipList) randomLevel() int {
	lvl := 1
	for rand.Float64() < p && lvl < maxLevel {
		lvl++
	}
	return lvl
}

func (sl *SkipList) Insert(val int) {
	update := make([]*Node, maxLevel)
	curr := sl.Head
	for i := maxLevel - 1; i >= 0; i-- {
		for curr.Next[i] != nil && curr.Next[i].Value < val {
			curr = curr.Next[i]
		}
		update[i] = curr
	}
	lvl := sl.randomLevel()
	newNode := &Node{Value: val, Next: make([]*Node, lvl)}
	for i := 0; i < lvl; i++ {
		newNode.Next[i] = update[i].Next[i]
		update[i].Next[i] = newNode
	}
}

func (sl *SkipList) Search(val int) bool {
	curr := sl.Head
	for i := maxLevel - 1; i >= 0; i-- {
		for curr.Next[i] != nil && curr.Next[i].Value < val {
			curr = curr.Next[i]
		}
	}
	curr = curr.Next[0]
	return curr != nil && curr.Value == val
}

func main() {
	sl := NewSkipList()
	sl.Insert(42)
	fmt.Println("Found 42:", sl.Search(42))
}
