package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single character",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "duplicate characters",
			input:    "aa",
			expected: []string{"aa"},
		},
		{
			name:     "two characters",
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:  "biro",
			input: "biro",
			expected: []string{
				"bior", "biro", "boir", "bori", "brio", "broi",
				"ibor", "ibro", "iobr", "iorb", "irbo", "irob",
				"obir", "obri", "oibr", "oirb", "orbi", "orib",
				"rbio", "rboi", "ribo", "riob", "robi", "roib",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateAnagrams(tt.input)

			// Sort both slices to ensure a consistent comparison
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("\nInput: %q\nGot: %v (len=%d)\n"+
					"Expected: %v (len=%d)",
					tt.input, result, len(result),
					tt.expected, len(tt.expected))
			}
		})
	}
}
