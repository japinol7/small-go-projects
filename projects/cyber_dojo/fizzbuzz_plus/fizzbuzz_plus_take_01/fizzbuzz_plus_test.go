package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"num 1", 1, "1"},
		{"num 2", 2, "2"},
		{"num 3", 3, "Fizz"},
		{"num 4", 4, "4"},
		{"num 5", 5, "Buzz"},
		{"num 6", 6, "Fizz"},
		{"num 10", 10, "Buzz"},
		{"num 13", 13, "Fizz"},
		{"num 15", 15, "FizzBuzz"},
		{"num 52", 52, "Buzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fizzbuzz(tt.input)
			if result != tt.expected {
				t.Errorf("Fizzbuzz(%d) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestFizzbuzzRange(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"range 1", 1, "1"},
		{"range 2", 2, "1\n2"},
		{"range 15", 15,
			"1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\n" +
				"Fizz\nFizz\n14\nFizzBuzz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := String(FizzbuzzRange(tt.input))
			if result != tt.expected {
				t.Errorf("ToString(FizzbuzzRange(%d)) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestFizzbuzzRangeUntilOneHundred(t *testing.T) {
	result := String(FizzbuzzRange(100))

	expected := strings.Join([]string{
		"1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\nFizz\n14\n",
		"FizzBuzz\n16\n17\nFizz\n19\nBuzz\nFizz\n22\nFizz\nFizz\nBuzz\n26\n",
		"Fizz\n28\n29\nFizzBuzz\nFizz\nFizz\nFizz\nFizz\nFizzBuzz\nFizz\n",
		"Fizz\nFizz\nFizz\nBuzz\n41\nFizz\nFizz\n44\nFizzBuzz\n46\n47\n",
		"Fizz\n49\nBuzz\nFizzBuzz\nBuzz\nFizzBuzz\nFizzBuzz\nBuzz\nBuzz\n",
		"FizzBuzz\nBuzz\nBuzz\nFizzBuzz\n61\n62\nFizz\n64\n",
		"Buzz\nFizz\n67\n68\nFizz\nBuzz\n71\nFizz\nFizz\n74\nFizzBuzz\n76\n",
		"77\nFizz\n79\nBuzz\nFizz\n82\nFizz\nFizz\nBuzz\n86\nFizz\n88\n89\n",
		"FizzBuzz\n91\n92\nFizz\n94\nBuzz\nFizz\n97\n98\nFizz\nBuzz",
	}, "")

	if result != expected {
		t.Errorf("FizzbuzzRange(100) produced incorrect result")
	}

	// Comment the following line to remove the output result
	fmt.Println(result)
}
