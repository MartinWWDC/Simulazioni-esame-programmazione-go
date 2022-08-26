/*
Scrivere un programma che legga da standard input il nome di un mese (in minuscolo).
Il programma deve stampare a video il numero di giorni di quel mese.
Si assume che l'anno non sia bisestile.

In particolare, si scriva una funzione numeroDiGiorni che,
dato in input il nome del mese (variabile di tipo stringa) restituisca il numero di giorni del mese (variabile di tipo intera).
La funzione deve usare il costrutto switch case.

INPUT:
$ go run mesi.go
Inserire mese: ottobre
Il numero di giorni di ottobre è 31

$ go run calcolatrice.go
Inserire mese: aprile
Il numero di giorni di aprile è 30

$ go run calcolatrice.go
Inserire mese: catalogna
Il nome del mese non è corretto.



*/

package main

import "fmt"

func numeroDiGiorni(mese string) (giorni int) {
	switch mese {
	case "gennaio", "marzo", "maggio", "luglio", "agosto", "ottobre", "dicembre":
		giorni = 31
	case "febbraio":
		giorni = 28
	case "settembre", "aprile", "giugno", "novembre":
		giorni = 30
	default:
		fmt.Println("mese non valido")
	}
	return
}

func main() {
	var mese string

	fmt.Print("Inserire mese:")
	fmt.Scanln(&mese)
	fmt.Println(numeroDiGiorni(mese))

}
