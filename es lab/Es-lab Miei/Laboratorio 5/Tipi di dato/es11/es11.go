package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Print("Inserisci il valore da troncare: ")
	fmt.Scan(&n)

	n = n * 100
	fmt.Print(math.Trunc(n) / 100)

}
