package anagrams

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to normalize the output for comparison.
// It sorts the inner slices and then the outer slice so we can use reflect.DeepEqual.
func normalize(input [][]string) {
	for _, sub := range input {
		sort.Strings(sub)
	}

	sort.Slice(input, func(i, j int) bool {
		first := input[i]
		second := input[j]
		if len(first) == 0 {
			return true
		}
		if len(second) == 0 {
			return false
		}

		return first[0] < second[0]
	})
}

func TestGroupAnagramsFrequency(t *testing.T) {
	// 1. Define Test Cases
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Standard Anagrams",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name:     "Single Character Strings",
			input:    []string{"a", "b", "a"},
			expected: [][]string{{"a", "a"}, {"b"}},
		},
		{
			name:     "Empty String in Slice",
			input:    []string{"", "b", ""},
			expected: [][]string{{"", ""}, {"b"}},
		},
		{
			name:     "No Anagrams",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "Empty Input Slice",
			input:    []string{},
			expected: [][]string{},
		},
	}

	// 2. Run Test Cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GroupAnagramsFrequency(tt.input)

			// Normalize both because map iteration order is random in Go
			normalize(actual)
			normalize(tt.expected)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("GroupAnagramsFrequency() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
