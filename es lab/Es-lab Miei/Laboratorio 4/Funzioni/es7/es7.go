package main

import (
	"fmt"
	"math"
)

func main() {
	var raggio float64

	fmt.Print("Inserisci il raggio del cerchio: ")
	fmt.Scan(&raggio)
	fmt.Print("Area del cerchio: ")

	fmt.Print(CalcolaArea(raggio))
	fmt.Println()
	fmt.Print("Circonferenza del cerchio: ")

	fmt.Print(CalcolaCirconferenza(raggio))

}

func CalcolaArea(raggio float64) float64 {
	return math.Pi * raggio * raggio
}

func CalcolaCirconferenza(raggio float64) float64 {
	return math.Pi * 2 * raggio
}
