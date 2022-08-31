package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Print("Inserisci il valore da arrotondare: ")
	fmt.Scan(&n)
	n = n * 100
	n = math.Round(n)
	fmt.Print(n / 100)

}
