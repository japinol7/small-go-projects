package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const startLetter = 'A'
const separator = " "

// Diamond represents a letter diamond pattern
type Diamond struct {
	letter  rune
	pattern []string
}

// NewDiamond creates a new diamond pattern for the given letter
func NewDiamond(letter string) *Diamond {
	d := &Diamond{}

	if utf8.RuneCountInString(letter) != 1 {
		return d
	}

	d.letter = unicode.ToUpper(rune(letter[0]))

	if !d.isInputValid() {
		return d
	}

	if d.letter == startLetter {
		d.pattern = []string{string(startLetter)}
		return d
	}

	d.createDiamondPattern()
	return d
}

// isInputValid checks if the input letter is valid for creating a diamond
func (d *Diamond) isInputValid() bool {
	if d.letter == 0 || d.letter < 'A' || d.letter > 'Z' {
		return false
	}
	return true
}

// addLinesUntilMiddleOfDiamond adds lines to the diamond pattern
// from A to the specified letter
func (d *Diamond) addLinesUntilMiddleOfDiamond(width int, lenToStartLetter int) {
	// Add the first line with just 'A'
	d.pattern = append(d.pattern, center(string(startLetter), width))

	// Add lines from 'B' to the specified letter
	for i := 1; i <= lenToStartLetter; i++ {
		ch := rune(int(startLetter) + i)
		spaces := 2*i - 1
		line := string(ch) + strings.Repeat(separator, spaces) + string(ch)
		d.pattern = append(d.pattern, center(line, width))
	}
}

// createDiamondPattern creates the diamond pattern
func (d *Diamond) createDiamondPattern() {
	lenToStartLetter := int(d.letter) - int(startLetter)
	width := 2*lenToStartLetter + 1

	d.addLinesUntilMiddleOfDiamond(width, lenToStartLetter)

	// Add mirrored lines to complete the diamond
	for i := len(d.pattern) - 2; i >= 0; i-- {
		d.pattern = append(d.pattern, d.pattern[i])
	}
}

// String returns the String representation of the diamond
func (d *Diamond) String() string {
	return strings.Join(d.pattern, "\n")
}

// center centers a String within a given width
func center(s string, width int) string {
	if len(s) >= width {
		return s
	}

	leftPadding := (width - len(s)) / 2
	rightPadding := width - len(s) - leftPadding

	return strings.Repeat(separator, leftPadding) + s +
		strings.Repeat(separator, rightPadding)
}
