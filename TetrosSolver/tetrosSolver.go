package tetrossolver

import (
	"fmt"
	"math"

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
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func MainSolver() {
	grid := makeGrid()
	solve(0, grid)
	printGrid(grid)
}

func solve(i int, grid [][]byte) {
	// ln, cl := 0, 0
	if i == len(*AllTetros) {
		return
	}
	
}
