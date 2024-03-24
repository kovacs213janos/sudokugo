package main

import (
	"fmt"
)

func euklalgo(a, b int) int {
	for b != 0 {
		fmt.Printf("a = %d , b = %d , maradek = %d \n", a, b, a%b)
		a, b = b, a%b
	}
	return a
}

func main() {
	a := 221
	b := 299
	lnko := euklalgo(a, b)
	fmt.Printf("A legnagyobb közös osztó: %d\n", lnko)
}
