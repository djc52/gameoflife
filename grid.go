// Grid contains the array of cells - live cells have curState = true
package main

import (
	"fmt"
)

// the game grid
type Grid struct {
	genIndex  int      // generation number
	maxCol    int      // number of columns in grid
	maxRow    int      // number of rows in grid
	gridCells [][]Cell // grid of cells
}

// getters
func (grid *Grid) GenIndex() int { return grid.genIndex }

// Grid "constructor""
func NewGrid(maxCol int, maxRow int) *Grid {
	grid := new(Grid)
	grid.maxCol = maxCol
	grid.maxRow = maxRow
	grid.gridCells = make([][]Cell, maxCol)
	for j := range grid.gridCells {
		grid.gridCells[j] = make([]Cell, maxRow)
	}
	return grid
}

// create a life at the given cell
func (grid *Grid) CreateLife(i int, j int) {
	grid.gridCells[i][j].curState = true
}

// print the grid to the console in a nice format
func (grid *Grid) PrintGrid() {
	var i, j int
	fmt.Print("\n\t+")
	for i = 0; i < grid.maxCol; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	for j = 0; j < grid.maxRow; j++ {
		fmt.Print("\n\t|")
		for i = 0; i < grid.maxCol; i++ {
			if grid.gridCells[i][j].curState == false {
				fmt.Print(" ")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Print("|")
	}
	fmt.Print("\n\t+")
	for i = 0; i < grid.maxCol; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
}

// utility to convert a bool to 0 or 1
func btou(b bool) int {
	if b {
		return 1
	}
	return 0
}

// computes one generation.  Each cell state in step n+1 is only dependent on the grid
// at step n.  So each cell has a curState and a nextState.  The nextState for all cells
// is calculated, and at the end it is swapped back to curState
func (grid *Grid) ComputeOneGeneration() {
	var i, j, im, ip, jm, jp int

	// copy curState to nextState for all cells
	for i = 0; i < grid.maxCol; i++ {
		for j = 0; j < grid.maxRow; j++ {
			grid.gridCells[i][j].nextState = grid.gridCells[i][j].curState
		}
	}

	var s00, s01, s02, s10, s12, s20, s21, s22 int

	// compute next state of each cell	
	for i = 0; i < grid.maxCol; i++ {
		// the computations are wrapped (so beings that 
		// move off the right re-enter on the left)
		// im and ip are the prior and next columns
		im = i - 1
		if i == 0 {
			im = grid.maxCol - 1
		}
		ip = i + 1
		if i == grid.maxCol-1 {
			ip = 0
		}
		for j = 0; j < grid.maxRow; j++ {
			// ditto for jm and jp
			jm = j - 1
			if j == 0 {
				jm = grid.maxRow - 1
			}
			jp = j + 1
			if j == grid.maxRow-1 {
				jp = 0
			}

			// convert surrounding grid states to ints to make rule testing easier
			s00 = btou(grid.gridCells[im][jm].curState)
			s10 = btou(grid.gridCells[i][jm].curState)
			s20 = btou(grid.gridCells[ip][jm].curState)
			s01 = btou(grid.gridCells[im][j].curState)
			s21 = btou(grid.gridCells[ip][j].curState)
			s02 = btou(grid.gridCells[im][jp].curState)
			s12 = btou(grid.gridCells[i][jp].curState)
			s22 = btou(grid.gridCells[ip][jp].curState)

			isAlive := grid.gridCells[i][j].curState

			sum := s00 + s10 + s20 + s01 + s21 + s02 + s12 + s22

			// if cell is live, and the number of neighboring live cells
			// is less than 2 or greater than 3, kill the cell
			if isAlive && (sum < 2 || sum > 3) {
				grid.gridCells[i][j].nextState = false
			}
			// if dead cell has exacty 3 neighbors, introduce life there
			if !isAlive && sum == 3 {
				grid.gridCells[i][j].nextState = true
			}

		}
	}

	//copy the next states back to the current states
	for i = 0; i < grid.maxCol; i++ {
		for j = 0; j < grid.maxRow; j++ {
			grid.gridCells[i][j].curState = grid.gridCells[i][j].nextState
		}
	}

	// increment generation count
	grid.genIndex++
}

// check if there is no life in grid
func (grid *Grid) IsEmpty() bool {
	for i := 0; i < grid.maxCol; i++ {
		for j := 0; j < grid.maxRow; j++ {
			if grid.gridCells[i][j].curState == true {
				return false
			}
		}
	}
	return true
}
