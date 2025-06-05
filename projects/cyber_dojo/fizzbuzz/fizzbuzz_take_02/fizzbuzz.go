package main

import (
	"fmt"
	"strings"
)

// Fizzbuzz converts a number to its FizzBuzz string representation
func Fizzbuzz(n int) string {
	var sb strings.Builder

	if n%3 == 0 {
		sb.WriteString("Fizz")
	}
	if n%5 == 0 {
		sb.WriteString("Buzz")
	}

	if sb.Len() == 0 {
		return fmt.Sprint(n)
	}
	return sb.String()
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
