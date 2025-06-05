package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LcdDigitsError represents an error in LCD digits generation
type LcdDigitsError struct {
	message string
}

func (e *LcdDigitsError) Error() string {
	return e.message
}

// NewLCDDigitsError creates a new LcdDigitsError
func NewLCDDigitsError(message string) *LcdDigitsError {
	return &LcdDigitsError{message: message}
}

// LcdDigits generates LCD-style digit representations
type LcdDigits struct{}

// Generate creates an LCD representation of a number
func (LcdDigits) Generate(num int, separator ...string) (string, error) {
	sep := Separator
	if len(separator) > 0 {
		sep = separator[0]
	}

	if num < 0 {
		return "", NewLCDDigitsError("Number must be non-negative")
	}

	// Handle the single digit case
	if num <= 9 {
		return fmt.Sprintf(NumToLcdDigitMapping[num], "\n", "\n", "\n"), nil
	}

	// Handle the multi-digit case
	digits, err := generateMultipleDigits(num, sep)
	if err != nil {
		return "", err
	}
	return strings.Join(digits, ""), nil
}

// generateMultipleDigits creates LCD representations for multi-digit numbers
func generateMultipleDigits(num int, separator string) ([]string, error) {
	numStr := strconv.Itoa(num)
	numLen := len(numStr)

	digitsTop := make([]string, 0, numLen)
	digitsMid := make([]string, 0, numLen)
	digitsBottom := make([]string, 0, numLen)

	for i, digitRune := range numStr {
		digit, err := strconv.Atoi(string(digitRune))
		if err != nil {
			return nil, NewLCDDigitsError("Invalid digit")
		}

		curSeparator := "\n"
		if i < numLen-1 {
			curSeparator = separator
		}

		digitTemplate := NumToLcdDigitMapping[digit]

		// Split the digit template into 3 parts
		topPart := fmt.Sprintf(digitTemplate[:5], curSeparator)
		midPart := fmt.Sprintf(digitTemplate[5:10], curSeparator)
		bottomPart := fmt.Sprintf(digitTemplate[10:15], curSeparator)

		digitsTop = append(digitsTop, topPart)
		digitsMid = append(digitsMid, midPart)
		digitsBottom = append(digitsBottom, bottomPart)
	}

	// Combine all parts
	result := make([]string, 0, len(digitsTop)+len(digitsMid)+len(digitsBottom))
	result = append(result, digitsTop...)
	result = append(result, digitsMid...)
	result = append(result, digitsBottom...)

	return result, nil
}
