package anagrams

import (
	"sort"
	"strings"
)

// GroupAnagramsFrequency groups strings by character frequency.
func GroupAnagramsFrequency(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for i := 0; i < len(s); i++ {
			// Assumes lowercase a-z as per typical interview constraints
			count[s[i]-'a']++
		}
		anagramMap[count] = append(anagramMap[count], s)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	// Map to store sorted string as key and original strings as a slice of values
	// Key: "aet" -> Value: ["eat", "tea", "ate"]
	anagramMap := make(map[string][]string)

	for _, s := range strs {
		// 1. Convert string to a slice of characters so we can sort it
		characters := strings.Split(s, "")
		sort.Strings(characters)

		// 2. Rejoin characters to create the sorted key
		key := strings.Join(characters, "")

		// 3. Append the original string to the corresponding key in the map
		anagramMap[key] = append(anagramMap[key], s)
	}

	// 4. Extract the values from the map into the final result slice
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
