package main

import (
	"fmt"
)

func main() {
	appName := "Print Diamond"
	fmt.Println("Start app", appName)

	// Examples of diamond patterns
	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {
		fmt.Printf("\nDiamond for letter %s:\n", letter)
		diamond := NewDiamond(letter)
		fmt.Println(diamond)
	}

	fmt.Println("\nEnd app", appName)
}
