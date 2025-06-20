package main

import (
	"testing"
)

func TestMineField_Resolve(t *testing.T) {
	tests := []struct {
		name     string
		board    string
		expected string
	}{
		{
			name: "test_resolve_1",
			board: "3 4\n" +
				"*...\n" +
				"..*.\n" +
				"....",
			expected: "*211\n" +
				"12*1\n" +
				"0111",
		},
		{
			name: "test_resolve_2",
			board: "5 4\n" +
				"*...\n" +
				"..*. \n" +
				"...*\n" +
				"....\n" +
				".*.. ",
			expected: "*211\n" +
				"12*2\n" +
				"012*\n" +
				"1121\n" +
				"1*10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := NewMineField(tt.board)
			result := mf.Resolve()
			if result != tt.expected {
				t.Errorf("MineField.Resolve() = %v, want %v", result, tt.expected)
			}
		})
	}
}
