package main

import (
	"fmt"
)

func main() {
	appName := "Quicksort for regular integers"
	fmt.Println("Start app", appName, "\n")

	numbers := []int{64, 34, -25, 14, 12, 22, 11, 90}
	fmt.Println("Original integers:", numbers)
	sortedNumbers := Quicksort(numbers)
	fmt.Println("Sorted integers:", sortedNumbers)
	fmt.Println("\nEnd app", appName)
}
