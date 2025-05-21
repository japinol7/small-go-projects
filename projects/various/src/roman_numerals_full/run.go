package main

import (
	"fmt"
	"log"
)

func main() {
	appName := "Roman Numerals Full"
	fmt.Println("Start app:", appName)

	// Examples of number to Roman numeral conversion
	nums := []int{1, 4, 9, 14, 42, 73, 99, 2023}
	fmt.Println("\nConverting numbers to Roman numerals:")
	for _, num := range nums {
		roman, err := ToRoman(num)
		if err != nil {
			log.Printf("Error converting %d: %v", num, err)
			continue
		}
		fmt.Printf("%d -> %s\n", num, roman)
	}

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
