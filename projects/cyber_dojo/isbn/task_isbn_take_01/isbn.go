package main

import (
	"strconv"
	"unicode"
)

// We represent the char X as -1
const xCharIntValue = -1

// validateISBN13 validates if a given string is a valid ISBN-13 code
func validateISBN13(isbnDigits []int, isbnDC int) string {
	// Multiply each digit alternately by 1 or 3 and sum the results
	checkDigit := 0
	for i, x := range isbnDigits {
		if (i+1)%2 == 0 {
			checkDigit += x * 3
		} else {
			checkDigit += x
		}
	}
	checkDigit %= 10
	checkDigit = (10 - checkDigit) % 10

	if checkDigit != isbnDC {
		return "false"
	}
	return "true"
}

// validateISBN10 validates if a given string is a valid ISBN-10 code
func validateISBN10(isbnDigits []int, isbnDC int) string {
	if isbnDC == xCharIntValue {
		isbnDC = 10
	}

	// Multiply each digit by its position number and sum the results
	checkDigit := 0
	for i, x := range isbnDigits {
		checkDigit += x * (i + 1)
	}
	checkDigit %= 11

	if checkDigit != isbnDC {
		return "false"
	}
	return "true"
}

// ValidateISBN validates if a given string is a valid ISBN-13 or ISBN-10 code
func ValidateISBN(isbn string) string {
	if len(isbn) == 0 {
		return "false"
	}

	// Check if all chars except the last one are numeric, '-', or space
	for i := 0; i < len(isbn)-1; i++ {
		ch := isbn[i]
		if !unicode.IsDigit(rune(ch)) && ch != '-' && ch != ' ' {
			return "false"
		}
	}

	// Extract digits for validation.
	// If the char is an X, extract it according to xCharIntValue
	var isbnNums []int
	for i := 0; i < len(isbn); i++ {
		if unicode.IsDigit(rune(isbn[i])) {
			digit, _ := strconv.Atoi(string(isbn[i]))
			isbnNums = append(isbnNums, digit)
		} else if isbn[i] == 'X' {
			isbnNums = append(isbnNums, xCharIntValue)
		}
	}

	isbnLen := len(isbnNums)
	isbnNumsWithoutDC := isbnNums[:isbnLen-1]
	isbnDC := isbnNums[isbnLen-1]

	if isbnLen == 13 {
		return validateISBN13(isbnNumsWithoutDC, isbnDC)
	}
	if isbnLen == 10 {
		return validateISBN10(isbnNumsWithoutDC, isbnDC)
	}
	return "false"
}
