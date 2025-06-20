package main

import (
	"fmt"
)

func main() {
	appName := "Mine field"
	fmt.Println("Start app", appName)

	// Example usage
	board := "3 4\n" +
		"*...\n" +
		"..*.\n" +
		"...."

	mf := NewMineField(board)
	result := mf.Resolve()

	fmt.Println("Input:")
	fmt.Println(board)
	fmt.Println("\nOutput:")
	fmt.Println(result)

	fmt.Println("End app", appName)
}
