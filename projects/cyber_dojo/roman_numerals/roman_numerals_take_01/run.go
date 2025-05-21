package main

import (
	"fmt"
	"log"
)

func main() {
	appName := "Roman Numerals"
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

	fmt.Println("\nEnd app:", appName)
}
