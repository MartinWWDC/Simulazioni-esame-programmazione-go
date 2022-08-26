/*
Scrivere un programma che legga da standard input un numero reale e ne stampi il valore troncato alla seconda cifra decimale.

Suggerimento:


Il valore troncato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:

Moltiplicare il numero reale per 100
Convertire il valore ottenuto al passo 1) in un valore di tipo int
Riconvertire il valore ottenuto al passo 2) in un valore di tipo float64
Dividere il valore ottenuto al passo 3) per 100


oppure:


Moltiplicare il numero reale per 100
Utilizzare la funzione math.Trunc del package math per troncare all'intero valore ottenuto al passo 1) (si utilizzi il comando go doc math.Trunc per capire come utilizzare la funzione)
Dividere il valore ottenuto al passo 2) per 100





Esempio d'esecuzione:

$ go run troncamento.go
Inserisci il valore da troncare: 10.34762
Valore troncato = 10.34

$ go run troncamento.go
Inserisci il valore da troncare: 8.34267
Valore troncato = 8.34
*/

package main

import "fmt"

func main() {

	/*
		Moltiplicare il numero reale per 100
		Convertire il valore ottenuto al passo 1) in un valore di tipo int
		Riconvertire il valore ottenuto al passo 2) in un valore di tipo float64
		Dividere il valore ottenuto al passo 3) per 100
	*/
	var n float64
	fmt.Scanln(&n)
	fmt.Println(n)

	n = n * 100
	fmt.Println(n)

	nint := int(n)
	n = float64(nint)

	fmt.Println(n)
	n = n / 100

	fmt.Println(n)

}
