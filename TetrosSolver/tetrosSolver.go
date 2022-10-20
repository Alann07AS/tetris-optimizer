package tetrossolver

import (
	"fmt"
	"math"

	tetrosclass "tetris/TetrosClass"
)

var AllTetros = tetrosclass.ParseAndGetAllTetros()

func makeGrid() [][]byte {
	size := int(math.Sqrt(float64(len(*AllTetros) * 4)))
	g := make([][]byte, size)
	for i := range g {
		g[i] = make([]byte, int(math.Sqrt(float64(len(*AllTetros)*4))))
	}
	return g
}

func printGrid(g [][]byte) {
	for iLn := range g {
		for _, char := range g[iLn] {
			tetrosclass.PrintIdAndColor(char)
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
	var actualGrid [][]byte
	copy(grid, actualGrid)
	if i == len(*AllTetros) {
		return
	}
	
}
