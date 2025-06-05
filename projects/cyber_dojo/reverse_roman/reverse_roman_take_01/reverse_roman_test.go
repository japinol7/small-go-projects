package main

import (
	"testing"
)

func TestFromRoman(t *testing.T) {
	testCases := []struct {
		roman    string
		expected int
		hasError bool
	}{
		{"I", 1, false},
		{"II", 2, false},
		{"III", 3, false},
		{"IV", 4, false},
		{"V", 5, false},
		{"IX", 9, false},
		{"X", 10, false},
		{"XL", 40, false},
		{"L", 50, false},
		{"LXXIII", 73, false},
		{"XC", 90, false},
		{"XCIII", 93, false},
		{"C", 100, false},
		{"CD", 400, false},
		{"D", 500, false},
		{"CM", 900, false},
		{"M", 1000, false},
		{"MCMLXXXIV", 1984, false},
		{"MMXXIII", 2023, false},
		{"MMMCMXCIX", 3999, false},
		{"", 0, true},          // Error case
		{"MMMM", 0, true},      // Error case (4000)
		{"ABC", 0, true},       // Error case (invalid chars)
		{"MMMCMXCIY", 0, true}, // Error case (invalid char Y)
	}

	for _, tc := range testCases {
		t.Run(tc.roman, func(t *testing.T) {
			result, err := FromRoman(tc.roman)
			if tc.hasError {
				if err == nil {
					t.Errorf("FromRoman(%s): expected an error, got nil",
						tc.roman)
				}
			} else {
				if err != nil {
					t.Errorf("FromRoman(%s): unexpected error: %v",
						tc.roman, err)
				}
				if result != tc.expected {
					t.Errorf("FromRoman(%s) = %d, expected %d",
						tc.roman, result, tc.expected)
				}
			}
		})
	}
}
