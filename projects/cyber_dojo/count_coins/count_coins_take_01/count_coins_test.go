package main

import (
	"fmt"
	"testing"
)

func TestChanges(t *testing.T) {
	tests := []struct {
		amount   int
		expected int
	}{
		{0, 0},
		{15, 6},
		{20, 9},
		{25, 13},
		{30, 18},
		{53, 49},
	}

	cc := &CountCoins{}

	for _, test := range tests {
		t.Run(fmt.Sprintf("amount=%d", test.amount), func(t *testing.T) {
			result := cc.Changes(test.amount)
			if result != test.expected {
				t.Errorf("Changes(%d) = %d, want %d",
					test.amount, result, test.expected)
			}
		})
	}
}

func TestChanges100CentsAndOutput(t *testing.T) {
	cc := &CountCoins{}
	result := cc.Changes(100)
	expected := 242

	if result != expected {
		t.Errorf("Changes(100) = %d, want %d", result, expected)
	}

	// Comment the following lines to remove the output result
	fmt.Println("Output: ")
	fmt.Println("How many ways are there to make change for a dollar ")
	fmt.Println("using these common coins? (1 dollar = 100 cents) ")
	fmt.Printf("Result: %d\n", result)
}
