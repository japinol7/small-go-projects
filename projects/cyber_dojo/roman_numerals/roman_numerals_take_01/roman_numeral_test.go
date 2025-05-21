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
