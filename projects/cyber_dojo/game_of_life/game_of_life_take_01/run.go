package main

import (
	"fmt"
)

func main() {
	appName := "Game of Life"
	fmt.Println("Start app", appName)

	gridPattern := []string{
		"........",
		"....*...",
		"...**...",
		".....*..",
	}

	game := NewGameOfLifeFromGrid(4, 8, gridPattern)

	err := game.SetCellRepresentation("*", ".", "")
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("End app", appName)
		return
	}

	generationsToRun := 4
	for i := 0; i < generationsToRun; i++ {
		fmt.Println("\nGeneration", i+1)
		fmt.Print(game.String())
		game.CalcNextGeneration()
	}

	fmt.Println("End app", appName)
}
