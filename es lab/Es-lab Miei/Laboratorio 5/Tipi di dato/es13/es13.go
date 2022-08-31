package main

import (
	"fmt"
	"math"
)

func main() {
	var n, molt float64
	var g int
	molt = 1
	fmt.Print("Inserisci il valore: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il numero di cifre dopo la virgola: ")
	fmt.Scan(&g)
	for i := 0; i < g; i++ {
		molt *= 10
	}
	fmt.Println(molt)
	n = n * molt
	fmt.Print("Valore troncato = ")
	fmt.Println(troncamento(int(molt), n))
	fmt.Print("Valore arrotondato = ")
	fmt.Println(arrotodnamento(int(molt), n))

}
func troncamento(g int, n float64) float64 {
	n = math.Trunc(n)
	return n / float64(g)
}
func arrotodnamento(g int, n float64) float64 {
	n = math.Round(n)
	return n / float64(g)
}
