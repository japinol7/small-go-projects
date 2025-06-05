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

// FromRoman converts a Roman numeral string to an integer
func FromRoman(roman string) (int, error) {
	if roman == "" {
		return 0, errors.New("empty Roman numeral string")
	}

	// Map for Roman numeral symbols to values
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	// Process from right to left
	for i := len(roman) - 1; i >= 0; i-- {
		char := rune(roman[i])
		value, exists := romanValues[char]
		if !exists {
			return 0, errors.New(
				"invalid Roman numeral character: " + string(char))
		}

		if value >= prevValue {
			result += value
		} else {
			result -= value
		}
		prevValue = value
	}

	// Check if the result is within the valid range for Roman numerals
	if result <= 0 || result > 3999 {
		return 0, errors.New("number out of range (1-3999)")
	}

	return result, nil
}
