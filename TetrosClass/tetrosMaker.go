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
	identifyTetros()
	printTetros()
	return &AllTetros
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

func identifyTetros() {
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
			for _, eachChar := range each {
				if eachChar == '#' {
					fmt.Print(string(AllTetros[i].CharId))
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
