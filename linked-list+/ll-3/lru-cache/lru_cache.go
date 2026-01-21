package main

import "fmt"

// Node represents a node in the Doubly Linked List
type Node struct {
	key, val   int
	prev, next *Node
}

// LRUCache structure
type LRUCache struct {
	capacity   int
	nodeMap    map[int]*Node
	head, tail *Node
}

// Constructor initializes the LRU Cache with a specific capacity
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		nodeMap:  make(map[int]*Node),
		head:     &Node{key: -1, val: -1}, // Dummy Head
		tail:     &Node{key: -1, val: -1}, // Dummy Tail
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get retrieves the value for a key and moves the node to the head
func (l *LRUCache) Get(key int) int {
	if node, ok := l.nodeMap[key]; ok {
		// Move to head: delete from current position and add after dummy head
		l.deleteNode(node)
		l.insertAtHead(node)
		return node.val
	}

	return -1
}

// Put adds or updates a key-value pair
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.nodeMap[key]; ok {
		// Existing key: remove and place same node at correct position with new value
		l.deleteNode(node)
		node.val = value
		l.insertAtHead(node)

		return
	}

	if len(l.nodeMap) == l.capacity {
		// Cache full: remove the LRU node (node before the dummy tail)
		lruNode := l.tail.prev
		delete(l.nodeMap, lruNode.key)
		l.deleteNode(lruNode)
	}

	// Add new node to DLL and Map
	newNode := &Node{key: key, val: value}
	l.insertAtHead(newNode)
	l.nodeMap[key] = newNode
}

// insertAtHead adds a new node right after the dummy head (Most Recently Used)
func (l *LRUCache) insertAtHead(newNode *Node) {
	currAfterHead := l.head.next
	l.head.next = newNode
	newNode.next = currAfterHead
	newNode.prev = l.head
}

// deleteNode removes a node from the Doubly Linked List
func (l *LRUCache) deleteNode(delNode *Node) {
	prevNode := delNode.prev
	nextNode := delNode.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)                     // Cache: [1]
	cache.Put(2, 2)                     // Cache: [2, 1]
	fmt.Println("Get 1:", cache.Get(1)) // Returns 1. Cache: [1, 2]

	cache.Put(3, 3)                     // Evicts key 2. Cache: [3, 1]
	fmt.Println("Get 2:", cache.Get(2)) // Returns -1 (not found)

	cache.Put(4, 4)                     // Evicts key 1. Cache: [4, 3]
	fmt.Println("Get 1:", cache.Get(1)) // Returns -1 (not found)
	fmt.Println("Get 3:", cache.Get(3)) // Returns 3. Cache: [3, 4]
	fmt.Println("Get 4:", cache.Get(4)) // Returns 4. Cache: [4, 3]
}
