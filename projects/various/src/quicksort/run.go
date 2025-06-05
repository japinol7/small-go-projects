package main

import (
	"fmt"
)

func main() {
	appName := "Quicksort"
	fmt.Println("Start app", appName, "\n")

	// Integer example
	numbers := []int{64, 34, -25, 14, 12, 22, 11, 90}
	fmt.Println("Original integers:", numbers)
	sortedNumbers := Quicksort(numbers)
	fmt.Println("Sorted integers:", sortedNumbers)

	// Float example
	floats := []float64{2.21, 3.44, 1.41, -0.75, 1.71, 0.51, 1.83}
	fmt.Println("\nOriginal floats:", floats)
	sortedFloats := Quicksort(floats)
	fmt.Println("Sorted floats:", sortedFloats)

	// String example
	words := []string{"banana", "eye", "reload", "holster", "arrow", "blue"}
	fmt.Println("\nOriginal strings:", words)
	sortedWords := Quicksort(words)
	fmt.Println("Sorted strings:", sortedWords)

	fmt.Println("\nEnd app", appName)
}
