To provide an even more exhaustive list, I have analyzed the "long-tail" of Glovo's interview history, including their move toward "Real-World Algorithm" scenarios. Glovo prefers questions that simulate their actual business logic (logistics, time-slots, and user sessions).

Here is the exhaustive categorized list of DSA questions asked at Glovo, including the specific "business context" they often use.

---

### 1. String & Array Manipulation (The "Warm-ups")

Glovo uses these to check for **clean code** and **test-driven development (TDD)** habits.

* **Anagram Detection:** Given two strings, return if they are anagrams. (Often extended: "Group all anagrams in a list of words").
* **First Unique Character:** Find the first non-repeating character in a stream of delivery IDs.
* **String Compression:** Implement `aabcccccaaa`  `a2b1c5a3`.
* **Palindrome Check:** Check if a string is a palindrome after removing non-alphanumeric characters.
* **Longest Common Prefix:** Finding common naming conventions in a list of product categories.

### 2. The "Logistics" & Sliding Window Patterns

These simulate moving windows of time or physical distance on a map.

* **Subarray Sum Equals K:** Find the number of continuous sections of a road where the customer density equals a target.
* **Sliding Window Maximum (The Ice Cream Truck):** Finding the busiest  consecutive sections of the beach.
* **Longest Substring Without Repeating Characters:** Finding the longest unique session of a user browsing the app.
* **Minimum Window Substring:** Finding the smallest stretch of a city that contains all required types of partner stores (e.g., one pharmacy, one grocery, one florist).

### 3. Stacks & Queues (The "Order Processing" Logic)

Since orders are often processed in specific sequences, stack/queue problems are a favorite.

* **Valid Parentheses:** The #1 most reported question.
* **Nearest Smaller Element:** Finding the closest cheaper alternative to a product.
* **Evaluate Reverse Polish Notation:** Evaluating complex tax/discount logic stored in postfix format.
* **Simplify Path:** Converting absolute paths in a server's file system (often asked in Senior roles).
* **Implement Queue using Stacks:** Testing your understanding of basic data structure trade-offs.

### 4. Two Pointers & Sorting (The "Optimization" Problems)

* **3Sum / 3Sum Closest:** Finding triplets of prices that hit a budget.
* **Container With Most Water:** Maximizing the "catchment area" for orders between two time-peaks.
* **Merge Sorted Intervals:** Essential for Glovo. "Given a set of delivery time slots, merge the overlapping ones."
* **Trapping Rain Water:** A high-level Two-Pointer/Stack problem often used for Senior candidates.

### 5. Linked Lists (The "History & Cache" Patterns)

* **LRU Cache:** Implement a cache for recently viewed restaurants. (Requires a HashMap + Doubly Linked List).
* **Reverse a Linked List:** A baseline technical check.
* **Detect Cycle in Linked List:** Checking for infinite loops in an order-state transition graph.

### 6. Trees & Graphs (The "Network & Category" Patterns)

Glovo uses these to represent "Category Trees" (Food > Pizza > Italian) or "Delivery Routes."

* **Lowest Common Ancestor:** Finding the closest shared category for two different items.
* **Level Order Traversal:** Printing all categories at a certain depth.
* **Course Schedule (Cycle Detection):** Determining if a set of delivery tasks has circular dependencies.
* **Number of Islands:** Finding clusters of busy delivery zones on a grid map.

---

### Summary of Glovo's "Interview Flavor"

If you get these questions, remember the **"Glovo Craftsmanship"** rules:

1. **Don't just code:** Explain the time complexity () and space complexity () before you type.
2. **Write Tests:** Before you say "I'm done," ask if you can write 2–3 test cases (Null input, Single element, Large input).
3. **Handle Edge Cases:** For Glovo,  or  are not just details—they are the test.

---

**Would you like me to provide a clean Go implementation for the "Merge Intervals" problem, as it is one of the most likely "Medium" questions you'll face for a Logistics/Backend role?**