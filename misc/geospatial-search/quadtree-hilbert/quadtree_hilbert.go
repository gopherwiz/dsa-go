/*
SPATIAL INDEXING STUDY NOTES: HILBERT CURVES & B-TREES (2026)
============================================================

1. HILBERT CURVE: THE SPATIAL "SNAKE"
   - Concept: A Space-Filling Curve that maps 2D coordinates (X, Y) into a 1D
     scalar value (Hilbert Index).
   - Locality Preservation: Unlike simple row-by-row scanning, the Hilbert curve
     ensures that points close to each other in 2D space remain close in 1D space.
   - Relationship to Quadtree: A Hilbert Curve is mathematically equivalent to
     traversing a Quadtree in a specific, recursive "U-shaped" pattern.

2. WHY USE HILBERT INSTEAD OF POINTER-BASED QUADTREE?
   - Database Friendly: Modern databases (B-Trees) are highly optimized for 1D
     sorting. By converting (X,Y) to 1D, we can use standard indexes.
   - Distributed Systems: It is easier to shard/partition a 1D range across multiple
     servers than it is to partition a recursive tree structure.
   - Cache Locality: Linearizing spatial data improves CPU cache hits when scanning
     nearby points.

3. PRODUCTION IMPLEMENTATION: GOOGLE S2 & POSTGRES
   - Google S2 Library: The industry standard (used by Uber, Foursquare, Google Maps).
     It projects the Earth onto a cube and applies a Hilbert Curve to each face.
   - Postgres (B-Tree + Hilbert): While PostGIS uses R-Trees, you can implement
     high-performance "point-in-radius" searches in vanilla Postgres by storing
     a Hilbert index in a standard B-Tree column.
   - Amazon DynamoDB: Often uses Hilbert/Z-Order curves to implement spatial
     searching on top of a NoSQL Key-Value store.



4. SEARCH MECHANISM: THE RANGE QUERY
   - Step 1 (Flatten): Convert the 2D Search Box (Bottom-Left, Top-Right) into
     start and end Hilbert indices.
   - Step 2 (Query): Perform a standard `BETWEEN start_h AND end_h` query.
   - Step 3 (Refine): Because the curve "snakes," a 1D range might include "false
     positives" (points outside the 2D box but inside the 1D range). A final
     check of the actual X, Y coordinates is required to filter these.

5. PERFORMANCE NOTES
   - Encoding (generateHilbert): O(log(GridSize)) — effectively constant time for
     64-bit integers.
   - Memory: Extremely low overhead. No pointers (unlike Quadtrees). Only stores
     the 1D index (int64).
   - Scaling: Handles billions of points easily because it relies on the
     underlying database's B-Tree efficiency.

6. GOLANG CODE SUMMARY (AS PER YOUR EXAMPLE)
   - Recursion: The `generateHilbert` function "walks" down the quadrants,
     rotating the orientation (xi, xj, yi, yj) to maintain the continuous curve.
   - Binary Search: `sort.Search` is used to find the first relevant H-Index in
     O(log N) time.
   - Precision: The `gridSize` determines the resolution. A gridSize of 8 means
     64 possible cells. For real-world GPS, 30-level Hilbert (S2) is common.
*/

/*
	QUADTREE WITH HILBERT (1D INDEXING)
	-----------------------------------
	- Concept: A Linear structure that "bakes" recursive logic into a
	single sortable number.
	- Storage Location: On-Disk (Database). Specifically designed to be
	stored in B-Trees (PostgreSQL/MySQL) or Sorted Strings (SSTables
	in Cassandra/BigTable).
	- Data Structure: A Sorted List or B-Tree Index. There are no pointers.
	Each record stores a uint64 (the Hilbert Index) as a column.
	- Disk Behavior: Extremely Efficient. B-Trees are built for "Block
	Storage." The database can find a spatial range with very few
	physical disk seeks by scanning the 1D Hilbert range.
	- Recursion Point: Happens during Encoding. You use recursion once to
	calculate the Index (X,Y -> H). After that, the recursion is
	discarded—the system only deals with a flat list of numbers.
*/

/*
	QUADTREE WITH HILBERT (1D INDEXING)
	-----------------------------------
	- Concept: A Linear structure that "bakes" recursive logic into a
	single sortable number.
	- Storage Location: On-Disk (Database). Specifically designed to be
	stored in B-Trees (PostgreSQL/MySQL) or Sorted Strings (SSTables
	in Cassandra/BigTable).
	- Data Structure: A Sorted List or B-Tree Index. There are no pointers.
	Each record stores a uint64 (the Hilbert Index) as a column.
	- Disk Behavior: Extremely Efficient. B-Trees are built for "Block
	Storage." The database can find a spatial range with very few
	physical disk seeks by scanning the 1D Hilbert range.
	- Recursion Point: Happens during Encoding. You use recursion once to
	calculate the Index (X,Y -> H). After that, the recursion is
	discarded—the system only deals with a flat list of numbers.
*/

package main

import (
	"fmt"
	"sort"
)

// Point and Cab structures (Same as before)
type Point struct{ X, Y int }
type Cab struct {
	ID     string
	HIndex int
}

// 1. The Recursive Hilbert Mapper
// x, y: coordinates
// s: current side length (must be power of 2)
// xi, xj, yi, yj: vectors defining the orientation of the curve
func generateHilbert(x, y, s, xi, xj, yi, yj int) int {
	if s <= 0 {
		return 0
	}

	// Determine which quadrant the point (x, y) falls into
	// This is effectively "walking down" the Quadtree
	quadX := x >= xi/2+xj/2
	quadY := y >= yi/2+yj/2

	// In an interview, explain that these rotations are what
	// make the Hilbert curve "stay close" in 1D space.
	if !quadX && !quadY { // Lower Left
		return generateHilbert(y, x, s/4, yi/2, yj/2, xi/2, xj/2)
	} else if !quadX && quadY { // Upper Left
		return s/4 + generateHilbert(x, y-yi/2, s/4, xi/2, xj/2, yi/2, yj/2)
	} else if quadX && quadY { // Upper Right
		return s/2 + generateHilbert(x-xi/2, y-yi/2, s/4, xi/2, xj/2, yi/2, yj/2)
	} else { // Lower Right
		return 3*s/4 + generateHilbert(xi/2-1-y, xj/2-1-x, s/4, -yi/2, -yj/2, -xi/2, -xj/2)
	}
}

// 2. The Search Logic (1D Range Search)
func searchCabs(cabs []Cab, startH, endH int) []Cab {
	var found []Cab
	// Binary search to find the start point in O(log N)
	idx := sort.Search(len(cabs), func(i int) bool {
		return cabs[i].HIndex >= startH
	})

	// Scan until we hit the end of the 1D range
	for i := idx; i < len(cabs) && cabs[i].HIndex <= endH; i++ {
		found = append(found, cabs[i])
	}
	return found
}

func main() {
	gridSize := 8 // 8x8 grid (2^3)
	maxIndex := gridSize * gridSize

	// Initial vectors for a standard Hilbert orientation
	xi, xj := gridSize, 0
	yi, yj := 0, gridSize

	// Create some cabs
	cabData := []struct {
		id   string
		x, y int
	}{
		{"Cab_A", 0, 0},
		{"Cab_B", 1, 1},
		{"Cab_C", 7, 7},
	}

	var cabList []Cab
	for _, c := range cabData {
		h := generateHilbert(c.x, c.y, maxIndex, xi, xj, yi, yj)
		cabList = append(cabList, Cab{c.id, h})
	}

	// Sort by Hilbert index to enable 1D range searching
	sort.Slice(cabList, func(i, j int) bool {
		return cabList[i].HIndex < cabList[j].HIndex
	})

	// Search for cabs in the bottom-left corner (0,0 to 2,2)
	// We convert the 2D box into a 1D range
	hStart := generateHilbert(0, 0, maxIndex, xi, xj, yi, yj)
	hEnd := generateHilbert(2, 2, maxIndex, xi, xj, yi, yj)

	results := searchCabs(cabList, hStart, hEnd)

	fmt.Printf("Searching from H-Index %d to %d\n", hStart, hEnd)
	for _, res := range results {
		fmt.Printf("Found: %s at H-Index %d\n", res.ID, res.HIndex)
	}
}
