/*
Scrivere una versione generalizzata dei programmi Troncamento e Arrotondamento che legga da standard input un intero n oltre al numero reale.
L'intero n specifica che il troncamento e l'arrotondamento devono avvenire alla n-esima cifra decimale.

Esempio d'esecuzione:
$ go run generalizzazione.go
Inserisci il valore: 10.34762
Inserisci il numero di cifre dopo la virgola: 4
Valore troncato = 10.3476
Valore arrotondato = 10.3476

$ go run generalizzazione.go
Inserisci il valore:  10.34762
Inserisci il numero di cifre dopo la virgola: 3
Valore troncato = 10.347
Valore arrotondato = 10.348

*/

package main

import (
	"fmt"
	"math"
)

func troncamento(n float64, divisore int) float64 {

	n = n * float64(divisore)

	nint := int(n)
	n = float64(nint)

	n = n / float64(divisore)

	return n
}
func arrotondamento(n float64, divisore int) float64 {

	//Moltiplicare il numero reale per 100
	n *= float64(divisore)

	//Utilizzare la funzione math.Round del package math per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando go doc math.Round per capire come utilizzare la funzione)
	n = math.Round(n)

	//Dividere il valore ottenuto al passo 2) per 100
	n /= float64(divisore)

	return n
}

func main() {
	var n float64
	var vir int
	div := 1

	fmt.Print("il valore:")
	fmt.Scanln(&n)

	fmt.Print("il numero di cifre dopo la virgola:")
	fmt.Scanln(&vir)

	for i := 0; i < vir; i++ {
		div *= 10
	}
	fmt.Println("Valore troncato = ", troncamento(n, div))
	fmt.Println("Valore arrotondamento = ", arrotondamento(n, div))

}
