package main

import (
	"strings"
	"testing"
)

// Chars used to display test LCD digits and the separator strings
const (
	TestCellH        = '_'
	TestCellV        = '|'
	TestCellO        = '.'
	TestSeparator    = " "
	TestSeparatorBig = "     "
)

const (
	DigitsRepr0 = "" +
		"._.\n" +
		"|.|\n" +
		"|_|\n"

	DigitsRepr1 = "" +
		"...\n" +
		"..|\n" +
		"..|\n"

	DigitsRepr2 = "" +
		"._.\n" +
		"._|\n" +
		"|_.\n"

	DigitsRepr3 = "" +
		"._.\n" +
		"._|\n" +
		"._|\n"

	DigitsRepr4 = "" +
		"...\n" +
		"|_|\n" +
		"..|\n"

	DigitsRepr5 = "" +
		"._.\n" +
		"|_.\n" +
		"._|\n"

	DigitsRepr6 = "" +
		"._.\n" +
		"|_.\n" +
		"|_|\n"

	DigitsRepr7 = "" +
		"._.\n" +
		"..|\n" +
		"..|\n"

	DigitsRepr8 = "" +
		"._.\n" +
		"|_|\n" +
		"|_|\n"

	DigitsRepr9 = "" +
		"._.\n" +
		"|_|\n" +
		"..|\n"

	DigitsRepr1234567890 = "" +
		"... ._. ._. ... ._. ._. ._. ._. ._. ._.\n" +
		"..| ._| ._| |_| |_. |_. ..| |_| |_| |.|\n" +
		"..| |_. ._| ..| ._| |_| ..| |_| ..| |_|\n"

	DigitsRepr1234567890SeparatorBig = "" +
		"...     ._.     ._.     ...     ._.     ._.     ._.     ._.     ._.     ._.\n" +
		"..|     ._|     ._|     |_|     |_.     |_.     ..|     |_|     |_|     |.|\n" +
		"..|     |_.     ._|     ..|     ._|     |_|     ..|     |_|     ..|     |_|\n"

	DigitsRepr910 = "" +
		"._. ... ._.\n" +
		"|_| ..| |.|\n" +
		"..| ..| |_|\n"
)

// replaceLCDDigitCells replaces cell chars with test chars
func replaceLCDDigitCells(digitCellStr string) string {
	return strings.NewReplacer(
		string(CellH), string(TestCellH),
		string(CellV), string(TestCellV),
		string(CellO), string(TestCellO),
	).Replace(digitCellStr)
}

// replaceLCDDigitSeparators replaces separators with test separators
func replaceLCDDigitSeparators(
	digitCellStr string, separator, testSeparator string,
) string {
	return strings.Replace(digitCellStr, separator, testSeparator, -1)
}

func TestGenerateLCDDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"0", 0, DigitsRepr0},
		{"1", 1, DigitsRepr1},
		{"2", 2, DigitsRepr2},
		{"3", 3, DigitsRepr3},
		{"4", 4, DigitsRepr4},
		{"5", 5, DigitsRepr5},
		{"6", 6, DigitsRepr6},
		{"7", 7, DigitsRepr7},
		{"8", 8, DigitsRepr8},
		{"9", 9, DigitsRepr9},
		{"all digits", 1234567890, DigitsRepr1234567890},
		{"910", 910, DigitsRepr910},
	}

	lcd := LcdDigits{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lcd.Generate(tt.input)
			if err != nil {
				t.Fatalf("Generate(%d) returned error: %v", tt.input, err)
			}

			result = replaceLCDDigitCells(result)
			result = replaceLCDDigitSeparators(
				result, Separator, TestSeparator)

			if result != tt.expected {
				t.Errorf("Generate(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGenerateLCDDigitsWithCustomSeparator(t *testing.T) {
	lcd := LcdDigits{}
	separator := "    "

	result, err := lcd.Generate(1234567890, separator)
	if err != nil {
		t.Fatalf("Generate(1234567890, %q) returned error: %v", separator, err)
	}

	result = replaceLCDDigitCells(result)
	result = replaceLCDDigitSeparators(result, separator, TestSeparatorBig)

	if result != DigitsRepr1234567890SeparatorBig {
		t.Errorf("Generate(1234567890, %q) produced incorrect result", separator)
	}
}

func TestNegativeNumShouldReturnError(t *testing.T) {
	lcd := LcdDigits{}

	_, err := lcd.Generate(-1)
	if err == nil {
		t.Error("Expected error for negative number, got nil")
	}
}
