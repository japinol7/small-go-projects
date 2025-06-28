package main

import (
	"sort"
)

// GenerateAnagrams generates all possible anagrams of the input text
// It returns a slice of strings containing all anagrams
func GenerateAnagrams(text string) []string {
	if text == "" {
		return []string{}
	}

	// Convert string to rune slice to handle UTF-8 characters properly
	runeText := []rune(text)
	var anagrams []string
	generatePermutations(runeText, 0, &anagrams)

	// Sort anagrams to match the expected output
	sort.Strings(anagrams)
	return anagrams
}

// generatePermutations generates all permutations of runes
// in text using Heap's algorithm
func generatePermutations(text []rune, start int, anagrams *[]string) {
	if start == len(text)-1 {
		*anagrams = append(*anagrams, string(text))
		return
	}

	// Generate permutations with current first character fixed
	generatePermutations(text, start+1, anagrams)

	// Generate permutations for each character as first character
	for i := start + 1; i < len(text); i++ {
		// Skip swapping if characters are the same to avoid duplicates
		if text[start] == text[i] {
			continue
		}

		// Swap characters
		text[start], text[i] = text[i], text[start]

		// Recurse with the new arrangement
		generatePermutations(text, start+1, anagrams)

		// Swap back to restore original arrangement
		text[start], text[i] = text[i], text[start]
	}
}
