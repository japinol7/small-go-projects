package main

import (
	"errors"
	"strings"
)

// ToRoman converts an integer to a Roman numeral string
func ToRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", errors.New("number out of range (1-3999)")
	}

	// Define Roman numeral symbols and their values
	values := []int{
		1000, 900, 500, 400, 100, 90, 50,
		40, 10, 9, 5, 4, 1,
	}
	numerals := []string{
		"M", "CM", "D", "CD", "C", "XC", "L",
		"XL", "X", "IX", "V", "IV", "I",
	}

	var result strings.Builder
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result.WriteString(numerals[i])
			num -= values[i]
		}
	}

	return result.String(), nil
}
