package main

import (
	"strings"
)

// ColumnAlignment represents the alignment type for text columns
type ColumnAlignment string

const (
	LeftAlignment   ColumnAlignment = "<"
	RightAlignment  ColumnAlignment = ">"
	CenterAlignment ColumnAlignment = "^"
	ColSep                          = "$"
	FillCh                          = " "
)

// AlignColumns aligns the columns in the given text
// according to the specified alignment
func AlignColumns(text string, alignment ColumnAlignment) string {
	if text == "" {
		return ""
	}

	// Split the text into lines and then into parts
	// by the column separator
	lines := strings.Split(text, "\n")
	parts := make([][]string, len(lines))

	for i, line := range lines {
		// Remove trailing the column separator if present
		// and split by the column separator
		line = strings.TrimSuffix(line, ColSep)
		parts[i] = strings.Split(line, ColSep)
	}

	// Find the max width for each column
	maxCols := 0
	for _, part := range parts {
		if len(part) > maxCols {
			maxCols = len(part)
		}
	}

	// Create a slice to store the max width for each column
	widths := make([]int, maxCols)

	// Calculate the max width for each column
	for colIdx := 0; colIdx < maxCols; colIdx++ {
		maxWidth := 0
		for _, part := range parts {
			if colIdx < len(part) && len(part[colIdx]) > maxWidth {
				maxWidth = len(part[colIdx])
			}
		}
		widths[colIdx] = maxWidth
	}

	// Format each line according to the alignment
	var result strings.Builder

	for _, part := range parts {
		for colIdx, word := range part {
			padding := widths[colIdx] - len(word)

			switch alignment {
			case LeftAlignment:
				result.WriteString(word)
				result.WriteString(strings.Repeat(FillCh, padding))
			case RightAlignment:
				result.WriteString(strings.Repeat(FillCh, padding))
				result.WriteString(word)
			case CenterAlignment:
				leftPad := padding / 2
				rightPad := padding - leftPad
				result.WriteString(strings.Repeat(FillCh, leftPad))
				result.WriteString(word)
				result.WriteString(strings.Repeat(FillCh, rightPad))
			}

			// Add space between columns, but not after the last one
			if colIdx < len(part)-1 {
				result.WriteString(FillCh)
			}
		}
	}

	return result.String()
}
