// Conway's Game of Life program
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// option 1 - run a sample
	// option 2 - user supplies grid width, height, and initial locations of beings
	option := showIntroAndGetOption()

	var width, height int      // width and height of grid
	var xCoords, yCoords []int // x and y coords of initial beings

	if option == 1 {
		// hardcoded sample input
		width = 40
		height = 20
		xCoords = make([]int, 8)
		yCoords = make([]int, 8)
		xCoords = append(xCoords, 5, 6, 6, 5, 4, 10, 10, 10)
		yCoords = append(yCoords, 5, 6, 7, 7, 7, 10, 11, 12)
	} else {
		// ask user to supply
		width, height = askWidthHeight()
		xCoords, yCoords = askCells()
	}

	gameGrid := NewGrid(width, height)

	for i, _ := range xCoords {
		gameGrid.CreateLife(xCoords[i], yCoords[i])
	}

	// run 100 generations and print each to console
	for ig := 0; ig < 100; ig++ {
		gameGrid.PrintGrid()
		gameGrid.ComputeOneGeneration()
		if gameGrid.IsEmpty() {
			fmt.Printf("\n\nAll life died after %d generations\n\n", gameGrid.GenIndex())
			fmt.Printf("\n\n\t**** Thanks for playing!!!  Goodbye. ****\n\n")
			os.Exit(1)
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\n\nLife is still going after 100 generations\n\n")
	fmt.Printf("\n\n\t**** Thanks for playing!!!  Goodbye. ****\n\n")
}

// display game description and ask user for option
func showIntroAndGetOption() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\n")
	fmt.Printf("\t\t****** Welcome to Conway's Game of Life *****\n\n")
	fmt.Printf("\tConway's Game of Life is a simulated world in which beings\n")
	fmt.Printf("\tplaced on a 2 dimensional grid live, die and are born based\n")
	fmt.Printf("\tupon simple adjacency rules.  It is a zero player game in which\n")
	fmt.Printf("\tthe initial configuration determines all subsequent generations.\n")
	fmt.Printf("\tIn this program, all emulations run for 100 generations in the\n")
	fmt.Printf("\tconsole and then the program exits.\n\n")
	fmt.Printf("\tSee Wikipedia \"Conway's Game of Life\" for more information\n\n\n\n")

	fmt.Printf("To run a sample, enter 1.\n")
	fmt.Printf("To set up and run your own emulation - enter 2.\n")
	fmt.Printf("If at any time you want to exit - enter q.\n\n")
	fmt.Printf("Enter choice: ")

	var option int

	goodRead := false
	for goodRead == false {
		str, err := reader.ReadString('\n')

		// if there is an error on read - exit - since we should always be able to read string
		if err != nil {
			fmt.Printf("\n\nBadRead - exiting\n\n")
			os.Exit(1)
			return 0
		}

		// if the user enters "q" say goodbye and exit
		if str == "q" || str == "q\n" {
			fmt.Printf("\n\n\t\t   ***** Goodbye *****\n\n")
			os.Exit(1)
			return 0
		}

		// parse the string
		num, err2 := fmt.Sscan(str, &option)

		// see what option the user entered
		if err2 != nil || num != 1 || (option != 1 && option != 2) {
			fmt.Printf("\nBad Input, try again (or type q to quit)\n")
		} else {
			goodRead = true
		}
	}
	return option
}

// ask the user for the width and height of the game
func askWidthHeight() (int, int) {
	reader := bufio.NewReader(os.Stdin)

	var width, height int

	goodRead := false
	for goodRead == false {
		fmt.Printf("\nEnter the grid width and height and hit enter. (e.g, 60 30)")
		fmt.Printf("\nValues should be between 10 and 100.")
		fmt.Printf("\nEnter Values: ")

		str, err := reader.ReadString('\n')

		// if there is an error on read - exit - since we should always be able to read string
		if err != nil {
			fmt.Printf("\n\nBadRead - exiting\n\n")
			os.Exit(1)
			return 0, 0
		}

		// if the user enters "q" say goodbye and exit
		if str == "q" || str == "q\n" {
			fmt.Printf("\n\n\t\t   ***** Goodbye *****\n\n")
			os.Exit(1)
			return 0, 0
		}

		// parse the string
		num, err2 := fmt.Sscan(str, &width, &height)

		// if it is not two integers, allow for loop to ask again
		if err2 != nil || num != 2 {
			fmt.Printf("\nBad Input, try again (or type q to quit)\n")
		} else {
			goodRead = true
		}
	}

	// make sure width and height are reasonable
	if width > 100 {
		width = 100
	}
	if width < 10 {
		width = 10
	}
	if height > 100 {
		height = 100
	}
	if height < 10 {
		height = 10
	}

	return width, height
}

// ask the user which cells have a life in them
func askCells() ([]int, []int) {

	reader := bufio.NewReader(os.Stdin)
	xCoords := make([]int, 0, 50)
	yCoords := make([]int, 0, 50)
	goodRead := false
	for goodRead == false {
		fmt.Printf("\nEnter the cells that are alive x1 y1 x2 y2 ... (e.g., 20 20 21 20 22 20)")
		fmt.Printf("\nEnter Values: ")
		str, err := reader.ReadString('\n')

		// if there is an error on read - exit - since we should always be able to read string
		if err != nil {
			fmt.Printf("\n\nBadRead - exiting\n\n")
			os.Exit(1)
			return nil, nil
		}

		// if the user enters "q" say goodbye and exit
		if str == "q" || str == "q\n" {
			fmt.Printf("\n\n\t\t   ***** Goodbye *****\n\n")
			os.Exit(1)
			return nil, nil
		}

		// parse the string
		subStrings := strings.Fields(str)

		fmt.Printf("\n\nnumcoords = %d\n\n", len(subStrings))
		if len(subStrings) == 0 {
			fmt.Printf("\nBad Input, try again (or type q to quit)\n")
		} else if len(subStrings)%2 != 0 {
			fmt.Printf("\nMust enter an even number of coordinates, try again\n")
		} else {
			coordValue := 0
			var err3 error
			isX := true
			for _, subString := range subStrings {
				coordValue, err3 = strconv.Atoi(subString)
				if err3 != nil {
					fmt.Printf("\nBad Input, try again (or type q to quit)\n")
					return nil, nil
				}
				if isX {
					xCoords = append(xCoords, coordValue)
					isX = false
				} else {
					yCoords = append(yCoords, coordValue)
					isX = true
				}
			}
			goodRead = true
		}
	}

	return xCoords, yCoords
}
