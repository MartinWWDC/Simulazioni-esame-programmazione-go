/*
Scrivere un programma che legga da standard input un valore intero e due valori reali:

il primo valore è il seme (seed) s da utilizzare per inizializzare il generatore di numeri casuali;
il secondo ed il terzo valore sono il coefficiente angolare m e il termine noto q di una retta r: y = m*x + q sul piano cartesiano.
Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali px e py che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:

determinare se, a meno di una costante EPSILON = 1.0e-9, il punto sta sopra, sotto, o appartiene alla retta r;
stampare a video il relativo messaggio (come mostrato nell'Esempio di esecuzione).
I valori px e py devono essere compresi nell'intervallo [0, 10.0).


*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` di almeno una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/* La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
a `y` a meno di una costante EPSILON */
/* ÈXUgualeAY(5.0,4.999) -> false */
/* ÈXUgualeAY(5.0,4.9999999999) -> true */
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` di almeno una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}

func main() {

	var seed int64

	var m, q float64

	fmt.Println("S:")
	fmt.Scanln(&seed)
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)

	rand.Seed(seed)

	for i := 0; i < 10; i++ {
		x := rand.Float64() * 5.0
		y := rand.Float64() * 5.0

		fmt.Print("Punto (", x, ",", y, ") - ")
		if ÈXMaggioreDiY(y, m*x+q) {
			fmt.Println("Il punto sta sopra la retta")

		} else if ÈXMinoreDiY(y, m*x+q) {
			fmt.Println("Il punto sta sotto la retta")
		} else {
			fmt.Println("Il punto appartiene alla retta")
		}
	}

}
