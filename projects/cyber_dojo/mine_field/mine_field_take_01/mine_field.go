package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Cell represents a cell in the minefield
type Cell string

const (
	Empty Cell = "."
	Mine  Cell = "*"
)

// MineField represents a field with mines
type MineField struct {
	nRows    int
	nColumns int
	board    [][]string
}

// NewMineField creates a new minefield from the given board string
func NewMineField(board string) *MineField {
	lines := strings.Split(board, "\n")

	// Parse dimensions from the first line
	dimensions := strings.Split(lines[0], " ")
	nRows, _ := strconv.Atoi(dimensions[0])
	nColumns, _ := strconv.Atoi(dimensions[1])

	// Parse the board
	boardGrid := make([][]string, nRows)
	for i := 0; i < nRows; i++ {
		boardRow := make([]string, nColumns)
		for j := 0; j < nColumns; j++ {
			boardRow[j] = string(lines[i+1][j])
		}
		boardGrid[i] = boardRow
	}

	return &MineField{
		nRows:    nRows,
		nColumns: nColumns,
		board:    boardGrid,
	}
}

// countNeighbours counts the number of mine neighbors for a given cell
func (m *MineField) countNeighbours(nRow, nColumn int) int {
	count := 0

	for x := max(nColumn-1, 0); x < min(nColumn+2, m.nColumns); x++ {
		for y := max(nRow-1, 0); y < min(nRow+2, m.nRows); y++ {
			// Skip the cell itself and count only if it's a mine
			if (x != nColumn || y != nRow) && m.board[y][x] == string(Mine) {
				count++
			}
		}
	}

	return count
}

// Resolve calculates the hint field based on the minefield
func (m *MineField) Resolve() string {
	// Create a copy of the board for the result
	res := make([][]string, m.nRows)
	for i := 0; i < m.nRows; i++ {
		res[i] = make([]string, m.nColumns)
		copy(res[i], m.board[i])
	}

	// Calculate the hint field
	for nRow := 0; nRow < m.nRows; nRow++ {
		for nColumn := 0; nColumn < m.nColumns; nColumn++ {
			if m.board[nRow][nColumn] == string(Mine) {
				res[nRow][nColumn] = string(Mine)
				continue
			}
			res[nRow][nColumn] = fmt.Sprintf("%d",
				m.countNeighbours(nRow, nColumn))
		}
	}

	// Join the result into a string
	resultLines := make([]string, m.nRows)
	for i := 0; i < m.nRows; i++ {
		resultLines[i] = strings.Join(res[i], "")
	}

	return strings.Join(resultLines, "\n")
}
