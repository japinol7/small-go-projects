package main

import "strings"

var charsToExcludeLeft = map[rune]bool{
	'\\': true,
	'/':  true,
}
var wordsToKeep = []string{
	"tests", "test", "spec", "step",
}
var separators = map[rune]bool{
	'-': true,
	'_': true,
	'.': true,
}

func isSeparator(ch rune) bool {
	return separators[ch]
}

func keepWordOnTheRight(text string, wordToKeep string, leftMark, rightMark int) int {
	wordIdx := strings.Index(text, wordToKeep)
	if wordIdx != -1 {
		if 0 < wordIdx && wordIdx < rightMark {
			rightMark = wordIdx
			if rightMark > 0 && isSeparator(rune(text[rightMark-1])) {
				rightMark--
			}
			rightMark += leftMark
		}
	}
	return rightMark
}

func keepWordOnTheLeft(text string, wordToKeep string, leftMark int) int {
	wordIdx := strings.Index(text, wordToKeep)
	if wordIdx != -1 {
		if wordIdx >= leftMark {
			leftMark = wordIdx + len(wordToKeep)
			if leftMark < len(text) && isSeparator(rune(text[leftMark])) {
				leftMark++
			}
		}
	}
	return leftMark
}

func FilenameRange(filename string) []int {
	if filename == "" {
		return []int{}
	}

	name := strings.ToLower(filename)

	// Remove chars from the left to the last char to exclude
	leftMark := 0
	for i := len(name) - 1; i >= 0; i-- {
		if charsToExcludeLeft[rune(name[i])] {
			leftMark = i + 1
			break
		}
	}

	// Remove chars from the right of the first dot char
	rightMark := len(name)
	for i, ch := range name {
		if ch == '.' {
			rightMark = i
			break
		}
	}

	// Remove words to keep on the right and their separators
	nameTP := name[leftMark:]
	for _, word := range wordsToKeep {
		rightMark = keepWordOnTheRight(nameTP, word, leftMark, rightMark)
	}

	// Remove words to keep on the left and their separators
	nameTP = name[:rightMark]
	for _, word := range wordsToKeep {
		leftMark = keepWordOnTheLeft(nameTP, word, leftMark)
	}

	return []int{leftMark, rightMark}
}
