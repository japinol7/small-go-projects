package main

import (
	"errors"
)

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
