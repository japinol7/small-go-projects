// Package main implements Conway's Game of Life
package main

import (
	"errors"
	"strings"
)

const cellAliveDefault = "*"
const cellEmptyDefault = "."
const separatorDefault = ""

// GameOfLife represents the game state
type GameOfLife struct {
	rowsN               int
	colsN               int
	grid                [][]int
	initialPattern      []string
	isInitialGeneration bool
	cellAlive           string
	cellEmpty           string
	separator           string
}

// NewGameOfLife creates a new game
func NewGameOfLife(rowsN, colsN int) *GameOfLife {
	grid := make([][]int, rowsN)
	emptyPattern := make([]string, rowsN)
	for i := range grid {
		grid[i] = make([]int, colsN)
		emptyPattern[i] = strings.Repeat(cellEmptyDefault, colsN)
	}
	return &GameOfLife{
		colsN:               colsN,
		rowsN:               rowsN,
		grid:                grid,
		initialPattern:      emptyPattern,
		isInitialGeneration: true,
		cellAlive:           cellAliveDefault,
		cellEmpty:           cellEmptyDefault,
		separator:           separatorDefault,
	}
}

// NewGameOfLifeFromGrid creates a new game from string grid representation
func NewGameOfLifeFromGrid(height, width int, gridStr []string) *GameOfLife {
	grid := make([][]int, height)
	initialPattern := make([]string, height)
	copy(initialPattern, gridStr)

	// Create grid using the default cell representation
	for i := range grid {
		grid[i] = make([]int, width)
		for j, ch := range gridStr[i] {
			if string(ch) == cellAliveDefault {
				grid[i][j] = 1
			}
		}
	}
	return &GameOfLife{
		colsN:               width,
		rowsN:               height,
		grid:                grid,
		initialPattern:      initialPattern,
		isInitialGeneration: true,
		cellAlive:           cellAliveDefault,
		cellEmpty:           cellEmptyDefault,
		separator:           separatorDefault,
	}
}

// SetCellRepresentation sets custom cell representation
func (g *GameOfLife) SetCellRepresentation(alive, empty, separator string) error {
	if !g.isInitialGeneration {
		return errors.New("cell representation can only be set in initial generation")
	}

	if len(alive) > 1 {
		return errors.New("alive cell representation must be a single character")
	}
	if len(empty) > 1 {
		return errors.New("empty cell representation must be a single character")
	}

	g.cellAlive = alive
	g.cellEmpty = empty
	g.separator = separator

	// Recreate grid using custom cell representation
	for i := range g.grid {
		for j := range g.grid[i] {
			if string(g.initialPattern[i][j]) == g.cellAlive {
				g.grid[i][j] = 1
			} else {
				g.grid[i][j] = 0
			}
		}
	}
	return nil
}

// SetCell sets the state of a cell at the given coordinates
func (g *GameOfLife) SetCell(x, y, state int) {
	g.grid[y][x] = state
}

// GetCell returns the state of a cell at the given coordinates
func (g *GameOfLife) GetCell(x, y int) int {
	return g.grid[y][x]
}

// CountNeighbors counts the number of live neighbors for a cell
func (g *GameOfLife) CountNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < g.colsN && ny >= 0 && ny < g.rowsN {
				count += g.grid[ny][nx]
			}
		}
	}
	return count
}

// CalcNextGeneration advances the game by one generation
func (g *GameOfLife) CalcNextGeneration() {
	newGrid := make([][]int, g.rowsN)
	for i := range newGrid {
		newGrid[i] = make([]int, g.colsN)
	}

	// Enforce the game rules to determine life and death
	for y := 0; y < g.rowsN; y++ {
		for x := 0; x < g.colsN; x++ {
			neighbors := g.CountNeighbors(x, y)
			if g.grid[y][x] == 1 {
				// Cell is alive
				if neighbors < 2 || neighbors > 3 {
					// Cell dies by isolation or by overpopulation
					newGrid[y][x] = 0
				} else {
					// Cell survives
					newGrid[y][x] = 1
				}
			} else {
				// Cell is dead
				if neighbors == 3 {
					// Cell becomes alive
					newGrid[y][x] = 1
				} else {
					// Cell stays dead
					newGrid[y][x] = 0
				}
			}
		}
	}
	g.grid = newGrid
	g.isInitialGeneration = false
}

// String displays the current grid state
func (g *GameOfLife) String() string {
	var result strings.Builder
	for _, row := range g.grid {
		for _, cell := range row {
			if cell == 1 {
				result.WriteString(g.cellAlive + g.separator)
				continue
			}
			result.WriteString(g.cellEmpty + g.separator)
		}
		result.WriteString("\n")
	}
	return result.String()
}
