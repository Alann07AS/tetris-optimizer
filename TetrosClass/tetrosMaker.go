package tetrosclass

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tetros struct {
	TableTetro [][]byte
	CharId     byte
	Color      string
}

var AllTetros []Tetros

func ParseAndGetAllTetros() *[]Tetros {
	parseTetros()
	identifyAndColorTetros()
	reformTetros()
	// printTetros()
	return &AllTetros
}

// Print Id And Color
func PrintIdAndColor(Id byte) {
	for _, each := range AllTetros {
		if each.CharId == Id {
			fmt.Print(each.Color)
			fmt.Print(string(Id))
		}
	}
}

// Put Tetro In Grid
func PutTetroInGrid(itetro int, grid [][]byte, ln, cl int) [][]byte {
	tetro := AllTetros[itetro].TableTetro
	sLn, sCl := ln, cl
	for ; ln < sLn+len(tetro); ln++ {
		for ; cl < sCl+len(tetro[0]); cl++ {
			if tetro[ln-sLn][cl-sCl] != '.' {
				grid[ln][cl] = tetro[ln-sLn][cl-sCl]
			}
		}
		cl = sCl
	}
	return grid
}

// Errase Tetro In Grid
func ErraseTetroInGrid(itetro int, grid [][]byte, ln, cl int) [][]byte {
	tetro := AllTetros[itetro].TableTetro
	sLn, sCl := ln, cl
	for ; ln < sLn+len(tetro); ln++ {
		for ; cl < sCl+len(tetro[0]); cl++ {
			if tetro[ln-sLn][cl-sCl] != '.' {
				grid[ln][cl] = 0
			}
		}
		cl = sCl
	}
	return grid
}

// check if is putable
func IsPuting(grid [][]byte, itetro, ln, cl int) bool {
	tetro := AllTetros[itetro].TableTetro
	sLn, sCl := ln, cl
	if len(tetro) > len(grid[ln:]) {
		return false
	}
	if len(tetro[0]) > len(grid[0][cl:]) {
		return false
	}
	for ; ln < sLn+len(tetro); ln++ {
		for ; cl < sCl+len(tetro[0]); cl++ {
			if tetro[ln-sLn][cl-sCl] != '.' && grid[ln][cl] != 0 {
				return false
			}
		}
		cl = sCl
	}
	return true
}

// Parse al Tettros in fille
func parseTetros() {
	// take txt fille
	txt, err := os.ReadFile("./TetrosTXT/" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	txtLn := strings.Fields(string(txt))
	// field [][]byte table of each tetros
	for i := 0; i <= len(txtLn)/4-1; i++ {
		tableTCompt := 0
		AllTetros = append(AllTetros, Tetros{})
		for _, each := range txtLn[i*4 : i*4+4] {
			AllTetros[i].TableTetro = append(AllTetros[i].TableTetro, []byte{})
			AllTetros[i].TableTetro[tableTCompt] = append(AllTetros[i].TableTetro[tableTCompt], []byte(each)...)
			tableTCompt++
		}
	}
	// Check if is true Teros
	for _, eachTeros := range AllTetros {
		CaseComp := 0
		ConectComp := 0
		for iLn, eachLn := range eachTeros.TableTetro {
			for iCase, eachCase := range eachLn {
				if eachCase == '#' {
					CaseComp++
					if iLn < len(eachTeros.TableTetro)-1 && eachTeros.TableTetro[iLn+1][iCase] == '#' {
						ConectComp++
					}
					if iLn > 0 && eachTeros.TableTetro[iLn-1][iCase] == '#' {
						ConectComp++
					}
					if iCase < len(eachTeros.TableTetro[iLn])-1 && eachTeros.TableTetro[iLn][iCase+1] == '#' {
						ConectComp++
					}
					if iCase > 0 && eachTeros.TableTetro[iLn][iCase-1] == '#' {
						ConectComp++
					}
				}
			}
		}
		if !(CaseComp == 4 && ConectComp >= 6) {
			log.Fatal(errors.New("bad tetros input"))
		}
	}
	// remove extra ln
	for iTeros, eachTeros := range AllTetros {
		var tempNewTable [][]byte
		for _, each := range eachTeros.TableTetro {
			if string(each) != "...." {
				tempNewTable = append(tempNewTable, each)
			}
		}
		AllTetros[iTeros].TableTetro = tempNewTable
	}
	// remove extra cl
	for iTeros, eachTeros := range AllTetros {
		iCl := []byte{}
		for i := 0; i < 4; i++ {
			clTemp := []byte{}
			for _, each := range eachTeros.TableTetro {
				clTemp = append(clTemp, each[i])
			}
			if !strings.Contains(string(clTemp), "#") {
				iCl = append(iCl, byte(i)+'0')
			}
		}
		for iLn, del := range eachTeros.TableTetro {
			newTemp := []byte{}
			for i, each := range del {
				if !strings.Contains(string(iCl), strconv.Itoa(i)) {
					newTemp = append(newTemp, each)
				}
			}
			AllTetros[iTeros].TableTetro[iLn] = newTemp
		}
	}
}

func identifyAndColorTetros() {
	colorRange := float64(600 / len(AllTetros))
	color := 0.0
	r, g, b := 0, 0, 0
	for i := range AllTetros {
		if color < 200 {
			if color < 100 {
				r = 255
				g = int(float64((color / 100) * 255))
			} else {
				r = int(float64((1 - ((color - 100) / 100)) * 255))
				g = 255
			}
			b = 0
		}
		if color >= 200 && color < 400 {
			if color < 300 {
				g = 255
				b = int(float64((color - 200) / 100 * 255))
			} else {
				g = int(float64((1 - ((color - 300) / 100)) * 255))
				b = 255
			}
			r = 0
		}
		if color >= 400 && color < 600 {
			if color < 500 {
				b = 255
				r = int(float64((color - 400) / 100 * 255))
			} else {
				b = int(float64((1 - ((color - 500) / 100)) * 255))
				r = 255
			}
			g = 0
		}
		color += colorRange
		AllTetros[i].CharId = 'A' + byte(i)
		AllTetros[i].Color = "\x1B[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
	}
}

func printTetros() {
	for i := range AllTetros {
		fmt.Print(AllTetros[i].Color)
		for _, each := range AllTetros[i].TableTetro {
			fmt.Println(string(each))
		}
	}
}

// replace # by IdChar
func reformTetros() {
	for i := range AllTetros {
		fmt.Print(AllTetros[i].Color)
		for iLn, each := range AllTetros[i].TableTetro {
			for iChar, eachChar := range each {
				if eachChar == '#' {
					AllTetros[i].TableTetro[iLn][iChar] = AllTetros[i].CharId
				}
			}
		}
	}
}
