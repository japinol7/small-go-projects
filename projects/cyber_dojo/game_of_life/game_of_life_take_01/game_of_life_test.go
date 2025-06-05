package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	game := NewGameOfLife(5, 5)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if game.GetCell(x, y) != 0 {
				t.Errorf("Expected cell (%d,%d) to be 0, got %d",
					x, y, game.GetCell(x, y))
			}
		}
	}
}

func TestSetAndGetCell(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.SetCell(1, 2, 1)
	if game.GetCell(1, 2) != 1 {
		t.Errorf("Expected cell (1,2) to be 1, got %d",
			game.GetCell(1, 2))
	}
}

func TestCountNeighbors(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.SetCell(1, 1, 1)
	game.SetCell(2, 1, 1)
	game.SetCell(1, 2, 1)

	neighbors1 := game.CountNeighbors(1, 1)
	if neighbors1 != 2 {
		t.Errorf("Expected cell (1,1) to have 2 neighbors, got %d",
			neighbors1)
	}

	neighbors2 := game.CountNeighbors(2, 2)
	if neighbors2 != 3 {
		t.Errorf("Expected cell (2,2) to have 3 neighbors, got %d",
			neighbors2)
	}
}

func TestStepUnderpopulation(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.SetCell(1, 1, 1)

	game.CalcNextGeneration()
	_ = game.SetCellRepresentation("*", ".", "")

	if game.GetCell(1, 1) != 0 {
		t.Errorf("Expected cell (1,1) to die from underpopulation, got %d",
			game.GetCell(1, 1))
	}
}

func TestStepOverpopulation(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.SetCell(1, 1, 1)
	game.SetCell(0, 0, 1)
	game.SetCell(0, 1, 1)
	game.SetCell(0, 2, 1)
	game.SetCell(1, 0, 1)

	game.CalcNextGeneration()
	_ = game.SetCellRepresentation("*", ".", "")

	if game.GetCell(1, 1) != 0 {
		t.Errorf("Expected cell (1,1) to die from overpopulation, got %d",
			game.GetCell(1, 1))
	}
}

func TestStepReproduction(t *testing.T) {
	input := []string{
		"**...",
		"*....",
		".....",
		".....",
		".....",
	}
	game := NewGameOfLifeFromGrid(5, 5, input)
	_ = game.SetCellRepresentation("*", ".", "")

	game.CalcNextGeneration()
	if game.GetCell(1, 1) != 1 {
		t.Errorf("Expected cell (1,1) to become alive from reproduction, got %d",
			game.GetCell(1, 1))
	}
}

func TestNextGenerationFromGrid1(t *testing.T) {
	input := []string{
		"........",
		"....*...",
		"...**...",
		".....*..",
	}
	expected := []string{
		"........",
		"...**...",
		"...***..",
		"....*...",
	}

	game := NewGameOfLifeFromGrid(4, 8, input)
	_ = game.SetCellRepresentation("*", ".", "")
	game.CalcNextGeneration()

	for y := 0; y < 4; y++ {
		for x := 0; x < 8; x++ {
			expectedState := 0
			if expected[y][x] == '*' {
				expectedState = 1
			}
			if game.GetCell(x, y) != expectedState {
				t.Errorf("At position (%d,%d): expected %d, got %d",
					x, y, expectedState, game.GetCell(x, y))
			}
		}
	}
}

func TestNextGenerationFromGrid2(t *testing.T) {
	input := []string{
		"........",
		"...**...",
		".*****..",
		"........",
		"........",
	}
	expected := []string{
		"........",
		".....*..",
		"..*..*..",
		"..***...",
		"........",
	}

	game := NewGameOfLifeFromGrid(5, 8, input)
	_ = game.SetCellRepresentation("*", ".", "")
	game.CalcNextGeneration()

	for y := 0; y < 4; y++ {
		for x := 0; x < 8; x++ {
			expectedState := 0
			if expected[y][x] == '*' {
				expectedState = 1
			}
			if game.GetCell(x, y) != expectedState {
				t.Errorf("At position (%d,%d): expected %d, got %d",
					x, y, expectedState, game.GetCell(x, y))
			}
		}
	}
}

func TestNextGenerationFromGridCustomCellRepr(t *testing.T) {
	input := []string{
		"--------",
		"---@@---",
		"-@@@@@--",
		"--------",
		"--------",
	}
	expected := []string{
		"--------",
		"-----@--",
		"--@--@--",
		"--@@@---",
		"--------",
	}

	game := NewGameOfLifeFromGrid(5, 8, input)
	_ = game.SetCellRepresentation("@", "-", " ")
	game.CalcNextGeneration()

	for y := 0; y < 4; y++ {
		for x := 0; x < 8; x++ {
			expectedState := 0
			if expected[y][x] == '@' {
				expectedState = 1
			}
			if game.GetCell(x, y) != expectedState {
				t.Errorf("At position (%d,%d): expected %d, got %d",
					x, y, expectedState, game.GetCell(x, y))
			}
		}
	}
}

func TestStepOverpopulationCustomCellRepr(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.SetCell(1, 1, 1)
	game.SetCell(0, 0, 1)
	game.SetCell(0, 1, 1)
	game.SetCell(0, 2, 1)
	game.SetCell(1, 0, 1)

	game.CalcNextGeneration()
	_ = game.SetCellRepresentation("@", "-", " ")

	if game.GetCell(1, 1) != 0 {
		t.Errorf("Expected cell (1,1) to die from overpopulation, got %d",
			game.GetCell(1, 1))
	}
}
