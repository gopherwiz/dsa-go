package main

import "fmt"

// Node represents a single character in the Trie
type Node struct {
	links [26]*Node
	flag  bool
}

// containsKey checks if a child node exists for the given character
func (n *Node) containsKey(ch byte) bool {
	return n.links[ch-'a'] != nil
}

// put creates a new reference node for a character
func (n *Node) put(ch byte, node *Node) {
	n.links[ch-'a'] = node
}

// get returns the reference node for a character
func (n *Node) get(ch byte) *Node {
	return n.links[ch-'a']
}

// setEnd marks the end of a word
func (n *Node) setEnd() {
	n.flag = true
}

// isEnd checks if a word ends at this node
func (n *Node) isEnd() bool {
	return n.flag
}

// Trie data structure
type Trie struct {
	root *Node
}

// Constructor initializes the Trie with a root node
func Constructor() Trie {
	return Trie{root: &Node{}}
}

// Insert adds a word into the trie
// Time Complexity: O(length of word) [00:26:05]
func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		if !node.containsKey(word[i]) {
			node.put(word[i], &Node{})
		}
		node = node.get(word[i])
	}
	node.setEnd()
}

// Search returns true if the word is in the trie
// Time Complexity: O(length of word) [00:29:29]
func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		if !node.containsKey(word[i]) {
			return false
		}
		node = node.get(word[i])
	}
	return node.isEnd()
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		if !node.containsKey(prefix[i]) {
			return false
		}
		node = node.get(prefix[i])
	}
	return true
}

func main() {
	trie := Constructor()

	// Example operations from the video [00:01:01]
	trie.Insert("apple")
	fmt.Println("Search 'apple':", trie.Search("apple"))     // Output: true
	fmt.Println("Search 'app':", trie.Search("app"))         // Output: false
	fmt.Println("StartsWith 'app':", trie.StartsWith("app")) // Output: true

	trie.Insert("apps")
	fmt.Println("Search 'apps':", trie.Search("apps")) // Output: true
}
