package main

import (
	"fmt"
	"log"
)

func main() {
	appName := "Reverse Roman Numerals"
	fmt.Println("Start app:", appName)

	// Examples of Roman numeral to number conversion
	romanNums := []string{
		"I", "IV", "IX", "XIV", "XLII", "LXXIII", "XCIX", "MMXXIII"}
	fmt.Println("\nConverting Roman numerals to numbers:")
	for _, roman := range romanNums {
		num, err := FromRoman(roman)
		if err != nil {
			log.Printf("Error converting %s: %v", roman, err)
			continue
		}
		fmt.Printf("%s -> %d\n", roman, num)
	}

	fmt.Println("\nEnd app:", appName)
}
