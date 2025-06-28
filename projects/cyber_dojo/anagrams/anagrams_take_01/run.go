package main

import (
	"fmt"
	"strings"
)

func main() {
	appName := "Anagrams"
	fmt.Println("Start app", appName)

	inputWords := []string{"a", "ab", "biro"}

	for _, inputWord := range inputWords {
		fmt.Printf("Anagrams for '%s':\n", inputWord)
		result := GenerateAnagrams(inputWord)
		if len(result) == 0 {
			fmt.Println("  None")
		} else {
			fmt.Printf("  %s\n", strings.Join(result, ", "))
		}
	}

	fmt.Println("End app", appName)
}
