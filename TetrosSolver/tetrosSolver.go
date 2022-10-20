package tetrossolver

import (
	"math"

	tetrosclass "tetris/TetrosClass"
)

var (
	AllTetros = tetrosclass.ParseAndGetAllTetros()
	Grid      = makeGrid()
)

func makeGrid() [][]byte {
	size := int(math.Sqrt(float64(len(*AllTetros) * 4)))
	g := make([][]byte, size)
	for i := range g {
		g[i] = make([]byte, int(math.Sqrt(float64(len(*AllTetros)*4))))
	}
	return g
}

func MainSolver() {
	
}
