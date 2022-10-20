package tetrossolver

import (
	"fmt"
	"math"
	"time"

	tetrosclass "tetris/TetrosClass"
)

var AllTetros = tetrosclass.ParseAndGetAllTetros()

func makeGrid() [][]byte {
	size := (math.Sqrt(float64(len(*AllTetros) * 4)))
	if size != float64(int(size)) {
		size = (size) + 1
	}
	g := make([][]byte, int(size))
	for i := range g {
		g[i] = make([]byte, int(size))
	}
	return g
}

func printGrid(g [][]byte) {
	for iLn := range g {
		for _, char := range g[iLn] {
			if char != 0 {
				tetrosclass.PrintIdAndColor(char)
				fmt.Print(" ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func MainSolver() {
	timeStart := time.Now()
	grid := makeGrid()
	for !solve(0, grid) {
		grid = EcraseGridSize(grid)
	}
	fmt.Println("Time taken :", time.Since(timeStart))
	printGrid(grid)
}

func solve(i int, grid [][]byte) bool {
	// var actualGrid [][]byte
	// copy(actualGrid, grid)
	if i == len(*AllTetros) {
		return true
	}
	for ln := 0; ln < len(grid); ln++ {
		for cl := 0; cl < len(grid[0]); cl++ {
			if tetrosclass.IsPuting(grid, i, ln, cl) {
				// fmt.Println("COUCOU", i, ln, cl)
				tetrosclass.PutTetroInGrid(i, grid, ln, cl)
				if solve(i+1, grid) {
					return true
				}
				// copy(grid, actualGrid)
				tetrosclass.ErraseTetroInGrid(i, grid, ln, cl)
			}
		}
	}
	return false
}

// ecrase grid size
func EcraseGridSize(grid [][]byte) [][]byte {
	for i := range grid {
		grid[i] = append(grid[i], 0)
	}
	grid = append(grid, []byte{})
	for range grid[0] {
		grid[len(grid[0])-1] = append(grid[len(grid[0])-1], 0)
	}
	return grid
}
