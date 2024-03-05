package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Sudoku [Size][Size]int

var Iteraciok int

const Size = 9

func Validalas(tabla Sudoku, sor, oszlop, szam int) bool {
	for i := 0; i < Size; i++ {
		if tabla[sor][i] == szam || tabla[i][oszlop] == szam {
			return false
		}
	}
	startSor, startOszlop := sor-sor%3, oszlop-oszlop%3
	for i := startSor; i < startSor+3; i++ {
		for j := startOszlop; j < startOszlop+3; j++ {
			if tabla[i][j] == szam {
				return false
			}
		}
	}
	return true
}

func Solver(tabla Sudoku) ([Size][Size]int, bool) {
	Iteraciok++
	for sor := 0; sor < Size; sor++ {
		for oszlop := 0; oszlop < Size; oszlop++ {
			if tabla[sor][oszlop] == 0 {
				for szam := 1; szam <= 9; szam++ {
					if Validalas(tabla, sor, oszlop, szam) {
						tabla[sor][oszlop] = szam
						Tabla_Kesz, elkeszult := Solver(tabla) //rekurziv hivas
						if elkeszult {
							return Tabla_Kesz, true // ilyenkor van kesz
						}
						tabla[sor][oszlop] = 0
					}
				}
				return tabla, false
			}
		}
	}
	return tabla, true // ilyenkor is kész
}

func kiiratasTabla(tabla Sudoku) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			fmt.Printf("%d ", tabla[i][j])
		}
		fmt.Println()
	}
}

func szambekeres() int {
	var szam int
	for {
		fmt.Print("Írj ide egy számot 1 és 81 között ")
		_, err := fmt.Scanln(&szam)
		if err != nil {
			fmt.Println("Invalid szám. Normálisan!")
			continue
		}

		if szam < 1 || szam > 81 {
			fmt.Println("A kapott szám nem 1 és 81 között van!")
			continue
		}
		break
	}
	return szam
}

func randomRemove(tabla Sudoku, csere int) Sudoku {
	for i := 0; i < csere; i++ {
		sor := rand.Intn(Size)
		oszlop := rand.Intn(Size)

		for tabla[sor][oszlop] == 0 {
			sor = rand.Intn(Size)
			oszlop = rand.Intn(Size)
		}
		tabla[sor][oszlop] = 0
	}
	return tabla
}

func main() {
	start := time.Now()

	Iteraciok = 0
	Tabla_Eredeti := Sudoku{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 7, 6},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	fmt.Println("Ez a kiindulási tábla:")
	kiiratasTabla(Tabla_Eredeti)
	fmt.Println("")
	randomSzam := szambekeres()
	fmt.Printf("A kapott szám: %d\n", randomSzam)
	fmt.Println(" Ennyit fogunk véletlenszerűen kivenni az eredeti táblából.")

	Tabla_Random := randomRemove(Tabla_Eredeti, randomSzam)
	fmt.Println("Ez a feladat:")
	kiiratasTabla(Tabla_Random)

	fmt.Println("")
	fmt.Println("Ez lett a megoldás:")
	fmt.Println("")

	result, v := Solver(Tabla_Random)
	if !v {
		fmt.Printf("Nem talált jó megoldást, iterációk száma: %d\n", Iteraciok)
	} else {
		kiiratasTabla(result)
		fmt.Printf("A megoldásig vezető iterációk száma: %d\n", Iteraciok)
	}
	elteltido := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Running time: %s\n", &elteltido)
}

//TODO iterációk számát globális változó helyett függvényparaméterként is át lehetne adni, pl.
// func Solver(tabla Sudoku) ([Size][Size]int, bool, int) {
//				return tabla, false, Iteraciok
//	megoldas, bool, Iteraciok := Solver(tabla)
