// Cell struct represents one cell in the Game Of Life
package main

// one cell in the grid
type Cell struct {
	curState  bool // state in current generation, true is alive
	nextState bool // state in next generation, true is alive
}

/*
// "constructor" to create new cell with input values
func NewCell(curState bool, nextState bool) *Cell {
	cell := new(Cell)
	cell.curState = curState
	cell.nextState = nextState
	return cell
}
*/
