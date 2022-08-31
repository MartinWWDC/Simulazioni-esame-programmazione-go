package main

import (
	"fmt"
	"math"
)

func main() {
	var n, radicen float64
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	radicen = math.Sqrt(n)
	fmt.Print(n, "uguale a ", radicen, " * ", radicen)

}
