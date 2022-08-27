package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	fmt.Print("Radice quadrata: ")
	fmt.Print(RadiceQuadrata(n))

}

func RadiceQuadrata(numero float64) (float64, bool) {
	if numero >= 0 {
		return math.Sqrt(numero), true
	} else {
		return 0, false
	}
}
