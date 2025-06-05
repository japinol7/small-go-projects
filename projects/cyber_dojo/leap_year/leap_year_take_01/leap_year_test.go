package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2001, false},
		{1996, true},
		{1900, false},
		{2000, true},
		{2023, false},
		{2024, true},
		{0, false},
	}

	for _, test := range tests {
		result := IsLeapYear(test.year)
		if result != test.expected {
			t.Errorf("IsLeapYear(%d) = %v; want %v",
				test.year, result, test.expected)
		}
	}
}
