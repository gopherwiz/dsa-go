/*
SPATIAL INDEXING STUDY NOTES: QUADTREES & POSTGRES (2026)
========================================================

1. QUADTREE ARCHITECTURE (THE RECURSIVE SPLIT)
   - Fundamental Rule: A 2D space is recursively divided into four quadrants
     (NorthWest, NorthEast, SouthWest, SouthEast).
   - Adaptive Density: Unlike a fixed grid, a Quadtree only splits when a region
     becomes "crowded" (reaches Capacity). This makes it highly efficient for
     unevenly distributed data (e.g., thousands of cabs in London, but only two in a desert).
   - Node Types:
     * Leaf Nodes: Store actual data points (Cabs/Points).
     * Internal Nodes: Store pointers to four child Quadtrees.

2. SEARCH MECHANISM: "PRUNING"
   - Spatial Intersect: When searching, if the search boundary does not overlap
     with a node's boundary, the entire branch is "pruned" (skipped).
   - Efficiency: This reduces search complexity from O(N) linear scan to O(log N).

3. PRODUCTION IMPLEMENTATION: POSTGRESQL (SP-GiST)
   - Structure: Postgres uses SP-GiST (Space-Partitioned Generalized Search Tree)
     to implement Quadtrees internally.
   - Use Case: Used for indexing geometric types (Points, Boxes, Polygons) in
     the standard 'cube' or 'earthdistance' modules.
   - Disk Optimization: While your Go code is in-memory, Postgres maps these
     recursive quadrants to disk pages, minimizing the number of times the
     hard drive is read.
   - Balancing: Postgres handles the "Deep Tree" problem where a single dense
     spot could create a tree so deep it slows down.

4. QUADTREE VS. R-TREE (POSTGIS)
   - Postgres/PostGIS actually defaults to R-Trees (Rectangle Trees) for most GIS.
   - Quadtree Logic: Space is divided into fixed quadrants regardless of object shape.
   - R-Tree Logic: Rectangles "wrap" around the actual objects and can overlap.
   - When to use Quadtree: Better for "Point" data and high-speed updates where
     simpler spatial partitioning is faster than calculating complex bounding boxes.

5. PERFORMANCE NOTES
   - Insertion: O(log N). May trigger a "Subdivide" which is a one-time CPU cost.
   - Memory: High pointer overhead. Every "Divided" node creates 4 new pointers.
   - Deletion: Can be complex because you must check if four children can be
     "collapsed" back into one parent if their total count falls below Capacity.

6. GOLANG CODE SUMMARY (AS PER YOUR PROVIDED EXAMPLE)
   - Boundary Check: Uses `Contains()` to filter which branch to descend.
   - Intersection: Uses `Intersects()` to stop searching irrelevant branches.
   - Re-distribution: During `Subdivide()`, existing points move from the parent
     slice down into the new children to maintain the "Leaf-only data" rule.
*/

/*
	NORMAL QUADTREE (POINTER-BASED)
	-------------------------------
	- Concept: A classic recursive tree that mimics a 4-way LinkedList.
	- Storage Location: Primarily In-Memory (RAM). It is highly volatile;
	storing it on disk is difficult because memory addresses (pointers)
	change every time the process restarts.
	- Data Structure: A "Linked" Tree. Each node is a struct containing a
	data slice and four explicit memory pointers (*NW, *NE, *SW, *SE)
	to heap-allocated child structs.
	- Disk Behavior: Poor. To save it, you must "Serialize" (flatten) the
	entire tree into a file, which is a slow O(N) operation.
	- Recursion Point: Happens during Traversal. To find a point, the CPU
	must literally "walk" through a chain of pointers from root to leaf.
*/

package main

import (
	"fmt"
)

// Point represents a coordinate in 2D space (Longitude/Latitude or X/Y)
type Point struct {
	X float64
	Y float64
}

// Cab represents our data item
type Cab struct {
	ID  string
	Pos Point
}

// Boundary defines the spatial area of a quadrant using Center and Half-Dimensions [00:06:23]
type Boundary struct {
	Center     Point
	HalfWidth  float64
	HalfHeight float64
}

// Contains checks if a point is within this boundary
func (b Boundary) Contains(p Point) bool {
	return p.X >= b.Center.X-b.HalfWidth &&
		p.X <= b.Center.X+b.HalfWidth &&
		p.Y >= b.Center.Y-b.HalfHeight &&
		p.Y <= b.Center.Y+b.HalfHeight
}

// Intersects checks if another boundary overlaps with this one (used for search)
func (b Boundary) Intersects(other Boundary) bool {
	return !(other.Center.X-other.HalfWidth > b.Center.X+b.HalfWidth ||
		other.Center.X+other.HalfWidth < b.Center.X-b.HalfWidth ||
		other.Center.Y-other.HalfHeight > b.Center.Y+b.HalfHeight ||
		other.Center.Y+other.HalfHeight < b.Center.Y-b.HalfHeight)
}

// QuadTree represents a node in the spatial tree [00:17:10]
type QuadTree struct {
	Boundary Boundary
	Capacity int
	Cabs     []Cab
	Divided  bool

	// Children nodes [00:07:04]
	NW *QuadTree
	NE *QuadTree
	SW *QuadTree
	SE *QuadTree
}

// NewQuadTree initializes a new QuadTree node
func NewQuadTree(boundary Boundary, capacity int) *QuadTree {
	return &QuadTree{
		Boundary: boundary,
		Capacity: capacity,
		Cabs:     make([]Cab, 0),
		Divided:  false,
	}
}

// Subdivide splits the quadrant into four smaller child quadrants [00:14:49]
func (qt *QuadTree) Subdivide() {
	x := qt.Boundary.Center.X
	y := qt.Boundary.Center.Y
	w := qt.Boundary.HalfWidth / 2
	h := qt.Boundary.HalfHeight / 2

	qt.NW = NewQuadTree(Boundary{Point{x - w, y + h}, w, h}, qt.Capacity)
	qt.NE = NewQuadTree(Boundary{Point{x + w, y + h}, w, h}, qt.Capacity)
	qt.SW = NewQuadTree(Boundary{Point{x - w, y - h}, w, h}, qt.Capacity)
	qt.SE = NewQuadTree(Boundary{Point{x + w, y - h}, w, h}, qt.Capacity)

	qt.Divided = true
}

// Insert adds a cab to the tree. If capacity is reached, it subdivides. [00:23:05]
func (qt *QuadTree) Insert(cab Cab) bool {
	if !qt.Boundary.Contains(cab.Pos) {
		return false
	}

	if len(qt.Cabs) < qt.Capacity && !qt.Divided {
		qt.Cabs = append(qt.Cabs, cab)
		return true
	}

	if !qt.Divided {
		qt.Subdivide()
		// Move existing cabs to children to maintain tree integrity
		oldCabs := qt.Cabs
		qt.Cabs = nil
		for _, c := range oldCabs {
			qt.insertIntoChildren(c)
		}
	}

	return qt.insertIntoChildren(cab)
}

func (qt *QuadTree) insertIntoChildren(cab Cab) bool {
	if qt.NW.Insert(cab) || qt.NE.Insert(cab) || qt.SW.Insert(cab) || qt.SE.Insert(cab) {
		return true
	}
	return false
}

// Search retrieves all cabs within a specific range [00:40:22]
func (qt *QuadTree) Search(rangeBound Boundary, found *[]Cab) {
	if !qt.Boundary.Intersects(rangeBound) {
		return
	}

	// Check cabs in this node
	for _, cab := range qt.Cabs {
		if rangeBound.Contains(cab.Pos) {
			*found = append(*found, cab)
		}
	}

	// Recurse into children
	if qt.Divided {
		qt.NW.Search(rangeBound, found)
		qt.NE.Search(rangeBound, found)
		qt.SW.Search(rangeBound, found)
		qt.SE.Search(rangeBound, found)
	}
}

func main() {
	// 1. Initialize Root: Center (0,0), Span 100x100, Capacity 1 cab per node [00:12:06]
	rootBoundary := Boundary{Center: Point{0, 0}, HalfWidth: 100, HalfHeight: 100}
	qt := NewQuadTree(rootBoundary, 1)

	// 2. Insert Cabs [00:14:05]
	cabs := []Cab{
		{ID: "Cab_A", Pos: Point{10, 10}},
		{ID: "Cab_B", Pos: Point{-20, -20}},
		{ID: "Cab_C", Pos: Point{15, 15}}, // Will trigger subdivision
		{ID: "Cab_D", Pos: Point{80, 80}},
	}

	for _, c := range cabs {
		qt.Insert(c)
	}

	// 3. Search for cabs near point (12, 12) within a 10x10 area
	searchArea := Boundary{Center: Point{12, 12}, HalfWidth: 10, HalfHeight: 10}
	var nearbyCabs []Cab
	qt.Search(searchArea, &nearbyCabs)

	fmt.Printf("Found %d cabs in search area:\n", len(nearbyCabs))
	for _, c := range nearbyCabs {
		fmt.Printf("- %s at (%.1f, %.1f)\n", c.ID, c.Pos.X, c.Pos.Y)
	}
}
