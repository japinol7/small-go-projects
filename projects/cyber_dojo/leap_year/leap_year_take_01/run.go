package main

import (
	"fmt"
)

func main() {
	appName := "Leap Year"
	fmt.Println("Start app:", appName)

	exampleYears := []int{
		2001, 1996, 1900, 2000, 2023, 2024,
	}

	for _, year := range exampleYears {
		result := IsLeapYear(year)
		fmt.Printf("Is %d a leap year: %t\n", year, result)
	}

	fmt.Println("End app:", appName)
}
