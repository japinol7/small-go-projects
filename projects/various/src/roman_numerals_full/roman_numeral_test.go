package main

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
		hasError bool
	}{
		{1, "I", false},
		{2, "II", false},
		{3, "III", false},
		{4, "IV", false},
		{5, "V", false},
		{9, "IX", false},
		{10, "X", false},
		{40, "XL", false},
		{50, "L", false},
		{73, "LXXIII", false},
		{90, "XC", false},
		{93, "XCIII", false},
		{100, "C", false},
		{400, "CD", false},
		{500, "D", false},
		{900, "CM", false},
		{1000, "M", false},
		{1984, "MCMLXXXIV", false},
		{2023, "MMXXIII", false},
		{3999, "MMMCMXCIX", false},
		{0, "", true},    // Error case
		{4000, "", true}, // Error case
		{-1, "", true},   // Error case
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := ToRoman(tc.num)
			if tc.hasError {
				if err == nil {
					t.Errorf("ToRoman(%d): expected an error, got nil",
						tc.num)
				}
			} else {
				if err != nil {
					t.Errorf("ToRoman(%d): unexpected error: %v",
						tc.num, err)
				}
				if result != tc.expected {
					t.Errorf("ToRoman(%d) = %s, expected %s",
						tc.num, result, tc.expected)
				}
			}
		})
	}
}

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

// Test round-trip conversion
func TestRoundTrip(t *testing.T) {
	// Test every 100 numbers to keep test runtime reasonable
	for i := 1; i <= 3999; i += 100 {
		roman, err := ToRoman(i)
		if err != nil {
			t.Errorf("ToRoman(%d) unexpected error: %v", i, err)
			continue
		}

		num, err := FromRoman(roman)
		if err != nil {
			t.Errorf("FromRoman(%s) unexpected error: %v", roman, err)
			continue
		}

		if num != i {
			t.Errorf("Round trip failed: %d -> %s -> %d", i, roman, num)
		}
	}
}
