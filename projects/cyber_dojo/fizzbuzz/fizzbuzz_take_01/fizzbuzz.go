package main

import (
	"fmt"
	"strings"
)

// Fizzbuzz converts a number to its FizzBuzz string representation
func Fizzbuzz(n int) string {
	res := ""

	if n%3 == 0 {
		res += "Fizz"
	}
	if n%5 == 0 {
		res += "Buzz"
	}

	if res == "" {
		return fmt.Sprint(n)
	}
	return res
}

// FizzbuzzRange generates FizzBuzz strings for numbers from 1 to num
func FizzbuzzRange(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		res[i-1] = Fizzbuzz(i)
	}
	return res
}

// String converts an slice of any values to a newline-separated string
func String(items []string) string {
	return strings.Join(items, "\n")
}
