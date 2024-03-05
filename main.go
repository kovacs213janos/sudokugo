package main

import (
	"fmt"
	"math/rand"
)

type Sudoku [Size][Size]int

type Cella struct {
	Sor, Oszlop int
}

const Size = 9

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
	Tabla_Eredeti := [Size][Size]int{
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
	fmt.Printf("A kapott szám:", randomSzam)
	fmt.Println(" Ennyit fogunk véletlenszerűen kivenni az eredeti táblából.")

	Tabla_Random := randomRemove(Tabla_Eredeti, randomSzam)

	kiiratasTabla(Tabla_Random)
}
