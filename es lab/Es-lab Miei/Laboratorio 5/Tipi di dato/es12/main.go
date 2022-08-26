/*

Scrivere un programma che legga da standard input un numero reale e ne stampi il valore arrotondato alla seconda cifra decimale.

Scrivere un programma che legga da standard input un numero reale e ne stampi il valore troncato alla seconda cifra decimale.

Moltiplicare il numero reale per 100
Utilizzare la funzione math.Round del package math per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando go doc math.Round per capire come utilizzare la funzione)
Dividere il valore ottenuto al passo 2) per 100




Esempio d'esecuzione:
$ go run arrotondamento.go
Inserisci il valore da arrotondare: 10.34762
Valore arrotondato = 10.35

$ go run arrotondamento.go
Inserisci il valore da arrotondare: 8.32467
Valore arrotondato = 8.32


*/

package main

import (
	"fmt"
	"math"
)

func main() {

	// Scrivere un programma che legga da standard input un numero reale e ne stampi il valore arrotondato alla seconda cifra decimale.
	var n float64
	fmt.Print("Inserisci il valore da arrotondare:")
	fmt.Scanln(&n)

	//Moltiplicare il numero reale per 100
	n *= 100

	//Utilizzare la funzione math.Round del package math per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando go doc math.Round per capire come utilizzare la funzione)
	n = math.Round(n)

	//Dividere il valore ottenuto al passo 2) per 100
	n /= 100
	fmt.Println("Valore arrotondato = ", n)

}
